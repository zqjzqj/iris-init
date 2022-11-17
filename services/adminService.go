package services

import (
	"jd-fxl/global"
	"jd-fxl/model"
	"jd-fxl/repositories"
	"jd-fxl/repositories/repoComm"
	"jd-fxl/repositories/repoInterface"
	"jd-fxl/sErr"
	"database/sql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"strings"
	"time"
)

func NewAdminService() AdminService {
	return AdminService{repo: repositories.NewAdminRepo()}
}

type AdminService struct {
	repo repoInterface.AdminRepo
}

func (admServ AdminService) LoginByPwd(username, pwd string) (model.Admin, error) {
	adm := admServ.repo.GetByUsername(username)
	if adm.ID == 0 {
		adm = admServ.repo.GetByPhone(username)
		if adm.ID == 0 {
			return adm, sErr.New("无效的用户名")
		}
	}
	if !adm.IsStatusYes() {
		return adm, sErr.New("该账号已被禁用，请联系管理员")
	}
	if !adm.CheckPwd(pwd) {
		return adm, sErr.New("无效的密码")
	}
	err := admServ.LoginSuccess(&adm)

	//刷新一下权限
	admServ.RefreshPermissions(&adm, true)
	return adm, err
}

//初始化adm账号 在初次运行迁移的时候调用
func (admServ AdminService) InitAdminAccount() (model.Admin, error) {
	pwd, salt := global.GeneratePassword("123456")
	admin := model.Admin{
		Username:      "admin",
		Status:        global.IsYes,
		RealName:      model.RoleAdminName,
		Avatar:        "",
		Sex:           0,
		Password:      pwd,
		Salt:          salt,
		Token:         sql.NullString{},
		TokenStatus:   global.IsNo,
		LastLoginTime: 0,
		RolesId:       model.RoleAdmin,
	}
	admin.ID = model.AdminRootId
	err := admServ.repo.Save(&admin)
	return admin, err
}

func (admServ AdminService) ListPage(ctx iris.Context) ([]map[string]interface{}, *global.Pager) {
	where := repoInterface.AdmSearchWhere{}
	_ = ctx.ReadQuery(&where)
	pager := global.NewPager(ctx)
	pager.SetTotal(admServ.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []map[string]interface{}{}, pager
	}
	where.SelectParams = repoComm.SelectFrom{
		Offset:  pager.Offset,
		Limit:   pager.Size,
		RetSize: pager.Size,
		OrderBy: []repoComm.OrderByParams{{
			Column: "ID",
			Desc:   false,
		}},
	}
	adm := admServ.repo.GetList(where)
	return admServ.ShowMapList(adm), pager
}

//获取一条数据根据ctx
//这里条件为ID 传入ctx是方便后续修改参数条件
func (admServ AdminService) GetItem(ctx iris.Context, _select ...string) model.Admin {
	return admServ.repo.GetByID(ctx.URLParamUint64("ID"), _select...)
}

func (admServ AdminService) GetByID(id uint64, _select ...string) model.Admin {
	return admServ.repo.GetByID(id, _select...)
}

//通过请求ctx编辑/新增一条数据
func (admServ AdminService) EditByCtx(ctx iris.Context, admID uint64) (model.Admin, error) {
	admValidator := AdminValidator{}
	err := ctx.ReadBody(&admValidator)
	if err != nil {
		return model.Admin{}, err
	}
	//用于固定adm
	if admID > 0 {
		admValidator.ID = admID
	}
	return admServ.EditByValidator(admValidator)
}

func (admServ AdminService) EditByValidator(admValidator AdminValidator) (model.Admin, error) {
	adm, err := admServ.GetAdmByValidate(admValidator)
	if err != nil {
		return adm, err
	}
	err = admServ.repo.Save(&adm)
	return adm, err
}

func (admServ AdminService) DeleteByCtx(ctx iris.Context) error {
	admId := ctx.PostValueInt64Default("ID", 0)
	if admId == model.AdminRootId {
		return sErr.New("不能删除默认账户")
	}
	_, err := admServ.repo.Delete("id", admId)
	return err
}

//只刷新拥有角色的名称
func (admServ AdminService) RefreshRolesName(adm *model.Admin) {
	if adm.IsRootRole() {
		adm.RolesName = []string{model.RoleAdminName}
		return
	}
	rolesId := strings.Split(adm.RolesId, ",")
	if len(rolesId) == 0 {
		adm.RolesName = []string{}
		return
	}
	rolesIdUint64 := global.StrArrToUintArr(rolesId)
	roleRepo := repositories.NewRolesRepo()
	roles := roleRepo.GetRolesByID(rolesIdUint64...)
	adm.RolesName = make([]string, 0, len(roles))
	for _, rr := range roles {
		adm.RolesName = append(adm.RolesName, rr.Name)
	}
}

//刷新拥有角色的权限和名称
func (admServ AdminService) RefreshPermissions(adm *model.Admin, force bool) {
	if !force {
		if adm.Permissions != nil && adm.RolesName != nil {
			return
		}
	}
	if adm.IsRootRole() {
		adm.Permissions = nil
		adm.RolesName = []string{model.RoleAdminName}
		return
	}
	rolesId := strings.Split(adm.RolesId, ",")
	if len(rolesId) == 0 {
		adm.Permissions = []string{}
		adm.RolesName = []string{}
		return
	}
	rolesIdUint64 := global.StrArrToUintArr(rolesId)
	roleRepo := repositories.NewRolesRepo()
	roles := roleRepo.GetRolesByID(rolesIdUint64...)
	if len(roles) == 0 {
		adm.Permissions = []string{}
		adm.RolesName = []string{}
		adm.RolesId = ""
		_ = admServ.repo.Save(adm, "RolesId")
		return
	}
	rolesIdUint64 = rolesIdUint64[:0]
	adm.RolesName = make([]string, 0, len(roles))
	for _, rr := range roles {
		adm.RolesName = append(adm.RolesName, rr.Name)
		rolesIdUint64 = append(rolesIdUint64, rr.ID)
	}
	rolePermRepo := repositories.NewRolesPermissionsRepo()
	adm.Permissions = rolePermRepo.GetPermissionsByRoles(rolesIdUint64...)
}

func (admServ AdminService) LoginSuccess(adm *model.Admin) error {
	now := time.Now()
	adm.Token = sql.NullString{
		String: global.GenerateToken(256),
		Valid:  true,
	}
	adm.TokenStatus = global.IsYes
	adm.Status = global.IsYes
	adm.LastLoginTime = now.Unix()
	return admServ.repo.Save(adm)
}

func (admServ AdminService) Logout(adm *model.Admin) error {
	adm.TokenStatus = global.IsNo
	return admServ.repo.Save(adm, "TokenStatus")
}

func (admServ AdminService) CheckPhoneValid(adm model.Admin) error {
	if !global.CheckPhone(adm.Phone.String) {
		return sErr.New("无效的手机号")
	}
	_adm := admServ.repo.GetByPhone(adm.Phone.String, "id")
	if _adm.ID > 0 && _adm.ID != adm.ID {
		return sErr.NewFmt("该手机号码%s已存在", adm.Phone.String)
	}
	return nil
}

func (admServ AdminService) CheckUsernameValid(adm model.Admin) error {
	if adm.Username == "" {
		return sErr.New("用户名不能为空")
	}
	_adm := admServ.repo.GetByUsername(adm.Username, "id")
	if _adm.ID > 0 && _adm.ID != adm.ID {
		return sErr.NewFmt("该用户名%s已存在", adm.Username)
	}
	return nil
}

func (admServ AdminService) CheckPermissionByRoute(r *router.Route, ctx iris.Context) bool {
	adm, ok := ctx.Values().Get("adm").(model.Admin)
	if !ok {
		return false
	}
	return admServ.HasPermission(adm, NewPermissionService().GeneratePermissionAuthIdentify(r.Method, r.Path))
}

func (admServ AdminService) HasPermission(adm model.Admin, permIdent string) bool {
	if adm.IsRootRole() {
		return true
	}
	admServ.RefreshPermissions(&adm, false)
	if len(adm.Permissions) == 0 {
		return false
	}
	if !global.InSlice(permIdent, adm.Permissions) {
		return false
	}
	return true
}

func (admServ AdminService) ShowMapList(adm []model.Admin) []map[string]interface{} {
	_adm := []map[string]interface{}{}
	for _, v := range adm {
		//刷新一下角色名称
		admServ.RefreshRolesName(&v)
		_adm = append(_adm, v.ShowMap())
	}
	return _adm
}

//验证参数 并返回到一个新的adm model
func (admServ AdminService) GetAdmByValidate(aValidator AdminValidator) (model.Admin, error) {
	err := aValidator.Validate()
	if err != nil {
		return model.Admin{}, err
	}
	var adm model.Admin
	if aValidator.ID > 0 {
		adm = admServ.repo.GetByID(aValidator.ID)
		if adm.ID == 0 {
			return adm, sErr.New("无效的ID")
		}
	} else {
		adm = model.Admin{}
	}
	adm.Username = aValidator.Username
	adm.Sex = aValidator.Sex
	if aValidator.Phone == "" {
		adm.Phone = sql.NullString{}
	} else {
		adm.Phone = sql.NullString{
			String: aValidator.Phone,
			Valid:  true,
		}
	}
	adm.QQ = aValidator.QQ
	adm.Status = aValidator.Status
	adm.RealName = aValidator.RealName
	if aValidator.Avatar != "" {
		adm.Avatar = aValidator.Avatar
	}
	if aValidator.Password != "" {
		adm.Password = aValidator.Password
		adm.Password, adm.Salt = global.GeneratePassword(adm.Password)
	}

	adm.Desc = aValidator.Desc

	//自身资料编辑不修改状态和角色
	if aValidator.Self != "1" {
		if aValidator.Status == global.IsYes {
			adm.Status = global.IsYes
		} else {
			adm.Status = global.IsNo
		}
		//超管不予其他角色重合
		for k := range aValidator.RolesId {
			if aValidator.RolesId[k] == model.RoleAdmin {
				aValidator.RolesId = []string{"*"}
				break
			}
		}
		adm.RolesId = strings.Join(aValidator.RolesId, ",")
	}

	//这里默认的系统超级管理员 不能被修改角色权限
	if adm.ID == model.AdminRootId {
		if !adm.IsRootRole() {
			if adm.RolesId != "" {
				return adm, sErr.New("系统默认管理账户权限不能被修改")
			}
			adm.RolesId = model.RoleAdmin
		}
		if !adm.IsStatusYes() {
			return adm, sErr.New("系统默认管理账户不能被禁用")
		}
	}
	err = admServ.CheckPhoneValid(adm)
	if err != nil {
		return adm, err
	}
	err = admServ.CheckUsernameValid(adm)
	if err != nil {
		return adm, err
	}
	return adm, nil
}

type AdminValidator struct {
	ID       uint64
	Username string `json:"Name" validate:"required" label:"账户名"`
	Phone    string `json:"Phone" validate:"required" label:"手机号码"`
	QQ       string `json:"QQ" label:"QQ号" validate:"max=20"`
	Status   uint8  `json:"Status" label:"状态"`
	RealName string `json:"RealName" validate:"required" label:"真实姓名"`
	Avatar   string `json:"Avatar" label:"头像"`
	Password string `json:"Password" label:"密码"`
	Sex      uint8  `label:"性别"`
	Desc     string `label:"简介"`
	RolesId  []string
	Self     string //用于请求的时候在控制器区分一下是否是编辑当前用户的资料
}

func (aValidator AdminValidator) Validate() error {
	if aValidator.Password != "" {
		if err := global.CheckPassword(aValidator.Password); err != nil {
			return err
		}
	}
	//这里表示是新增的时候
	if aValidator.ID == 0 && aValidator.Password == "" {
		return sErr.New("密码不能为空")
	}
	err := global.ValidateV9Struct(aValidator)

	if err != nil {
		return err
	}
	return nil
}
