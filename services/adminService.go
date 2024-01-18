package services

import (
	"database/sql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"iris-init/global"
	"iris-init/model"
	"iris-init/repositories"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
	"iris-init/sErr"
	"reflect"
	"time"
)

func NewAdminService() AdminService {
	return AdminService{repo: repositories.NewAdminRepo()}
}

func NewAdminServiceByOrm(orm any) AdminService {
	r := AdminService{repo: repositories.NewAdminRepo()}
	r.repo.SetOrm(orm)
	return r
}

func NewAdminServiceByRepo(repo repoInterface.AdminRepo) AdminService {
	return AdminService{repo: repo}
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
	admServ.RefreshPermissions(&adm, true, false)
	return adm, err
}

// 初始化adm账号 在初次运行迁移的时候调用
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
		RolesID:       model.RoleAdmin,
	}
	admin.ID = model.AdminRootID
	err := admServ.repo.Save(&admin)
	return admin, err
}

func (adminServ AdminService) ListPage(ctx iris.Context) ([]model.Admin, *global.Pager) {
	where := repoInterface.AdminSearchWhere{}
	_ = ctx.ReadQuery(&where)
	where.SelectParams = repoComm.SelectFrom{
		OrderBy: []repoComm.OrderByParams{
			{
				Column: "ID",
				Desc:   true,
			},
		},
		Preload: []repoComm.PreloadParams{
			{
				Query: "Organizer",
			},
		},
	}
	pager := global.NewPager(ctx)
	if pager.Size < 0 {
		return adminServ.repo.GetList(where), nil
	}
	pager.SetTotal(adminServ.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []model.Admin{}, pager
	}
	where.SelectParams.Offset = pager.Offset
	where.SelectParams.Limit = pager.Size
	where.SelectParams.RetSize = pager.Size
	return adminServ.repo.GetList(where), pager
}

func (adminServ AdminService) ListAvailable(_select ...string) []model.Admin {
	if len(_select) == 0 {
		_select = nil
	}
	return adminServ.repo.GetList(repoInterface.AdminSearchWhere{
		SelectParams: repoComm.SelectFrom{
			Select: _select,
		},
	})
}

func (adminServ AdminService) GetByWhere(where repoInterface.AdminSearchWhere) model.Admin {
	return adminServ.repo.GetByWhere(where)
}

func (adminServ AdminService) ScanByWhere(where repoInterface.AdminSearchWhere, dest any) error {
	return adminServ.repo.ScanByWhere(where, dest)
}

func (adminServ AdminService) ScanByOrWhere(dest any, where ...repoInterface.AdminSearchWhere) error {
	return adminServ.repo.ScanByOrWhere(dest, where...)
}

func (adminServ AdminService) UpdateByWhere(where repoInterface.AdminSearchWhere, data interface{}) (rowsAffected int64, err error) {
	return adminServ.repo.UpdateByWhere(where, data)
}

func (adminServ AdminService) ListByWhere(where repoInterface.AdminSearchWhere) []model.Admin {
	return adminServ.repo.GetList(where)
}

func (adminServ AdminService) TotalCount(where repoInterface.AdminSearchWhere) int64 {
	return adminServ.repo.GetTotalCount(where)
}

// 获取一条数据根据ctx
// 这里条件为ID 传入ctx是方便后续修改参数条件
func (admServ AdminService) GetItem(ctx iris.Context, _select ...string) model.Admin {
	return admServ.repo.GetByID(ctx.URLParamUint64("ID"), _select...)
}

func (admServ AdminService) GetByID(id uint64, _select ...string) model.Admin {
	return admServ.repo.GetByID(id, _select...)
}

func (adminServ AdminService) DeleteByPhone(phone string) error {
	_, err := adminServ.repo.DeleteByPhone(phone)
	return err
}

func (adminServ AdminService) GetByToken(token string, _select ...string) model.Admin {
	var v reflect.Value
	v = reflect.ValueOf(token)
	if !v.IsValid() { // 值不存在
		return model.Admin{}
	}
	return adminServ.repo.GetByToken(token, _select...)
}

func (adminServ AdminService) CheckTokenValid(admin model.Admin) error {
	var v reflect.Value
	v = reflect.ValueOf(admin.Token)
	if !v.IsValid() { // 值不存在
		return sErr.New("无效的用户登陆token")
	}
	_admin := adminServ.GetByToken(admin.Token.String, "id")
	if _admin.ID > 0 && admin.ID != _admin.ID {
		return sErr.NewFmt("用户登陆token.已存在:%s.", admin.Token)
	}
	return nil
}

func (adminServ AdminService) DeleteByToken(token string) error {
	_, err := adminServ.repo.DeleteByToken(token)
	return err
}

// 通过请求ctx编辑/新增一条数据
func (admServ AdminService) SaveByCtx(ctx iris.Context, admID uint64) (model.Admin, error) {
	admValidator := AdminValidator{}
	err := global.ScanValidatorByRequestPost(ctx, &admValidator)
	if err != nil {
		return model.Admin{}, err
	}
	//用于固定adm
	if admID > 0 && admValidator.Self == "1" {
		admValidator.ID = admID
	}
	return admServ.SaveByValidator(admValidator)
}

func (admServ AdminService) SaveByValidator(admValidator AdminValidator) (model.Admin, error) {
	adm, err := admServ.GetAdmByValidate(admValidator)
	if err != nil {
		return adm, err
	}
	err = admServ.Save(&adm)
	return adm, err
}

func (admServ AdminService) Save(admin *model.Admin, _select ...string) error {
	//其他类型的管理员需要save一下角色
	roleAdmServ := NewRolesAdminService()
	return admServ.repo.Transaction(func() error {
		err := admServ.repo.Save(admin, _select...)
		if err != nil {
			return err
		}
		return roleAdmServ.SaveByAdm(*admin)
	}, roleAdmServ.repo)
}

func (admServ AdminService) DeleteByCtx(ctx iris.Context) error {
	admId := uint64(ctx.PostValueInt64Default("ID", 0))
	if admId == model.AdminRootID {
		return sErr.New("不能删除默认账户")
	}
	_, err := admServ.repo.DeleteByID(admId)
	if err != nil {
		return err
	}
	roleAdmServ := NewRolesAdminServiceByOrm(admServ.repo.GetOrm())
	_ = roleAdmServ.DeleteByAdminID(admId)
	return nil
}

// 刷新拥有角色的权限和名称
// 超级管理员权限会返回nil 其他的没有权限则是空数组
// onlyName == true时 不会从db去刷新实际权限 所以 只在需要获取角色名称时使用
func (admServ AdminService) RefreshPermissions(adm *model.Admin, force, onlyRolesName bool) {
	if !force {
		if adm.Permissions != nil && adm.RolesName != nil {
			return
		}
	}
	RoleIDSlices := adm.RefreshRoleIDSlices()
	if len(RoleIDSlices) == 0 {
		adm.Permissions = []string{}
		adm.RolesName = []string{}
		return
	}

	//仅仅只有超级管理员的情况
	if len(RoleIDSlices) == 1 && RoleIDSlices[0] == model.RoleAdmin {
		adm.Permissions = []string{model.RoleAdmin}
		adm.RolesName = []string{model.RoleAdminName}
		return
	}
	RolesIDUint64 := global.StrArrToUintArr(RoleIDSlices)
	roles := NewRolesServiceByOrm(admServ.repo.GetOrm()).GetRolesByID(RolesIDUint64...)

	if len(roles) == 0 {
		//这里因为超级管理员是没有实际数据库ID的 是通过程序内部校验 所以多一步判断
		if adm.IsRootRole() {
			//去掉多余的权限ID
			adm.Permissions = []string{model.RoleAdmin}
			adm.RolesName = []string{model.RoleAdminName}
			adm.SetRoleID([]string{model.RoleAdmin})
			_ = admServ.Save(adm, "RolesID")
		} else {
			adm.Permissions = []string{}
			adm.RolesName = []string{}
		}
		return
	}

	adm.RolesName = make([]string, 0, len(roles)+1)
	if adm.IsRootRole() {
		adm.Permissions = []string{model.RoleAdmin}
		adm.RolesName = append(adm.RolesName, model.RoleAdminName)
		for _, rr := range roles {
			adm.RolesName = append(adm.RolesName, rr.Name)
		}
		return
	}
	RolesIDUint64 = RolesIDUint64[:0]
	for _, rr := range roles {
		adm.RolesName = append(adm.RolesName, rr.Name)
		RolesIDUint64 = append(RolesIDUint64, rr.ID)
	}

	//这里不进行db查询了
	if onlyRolesName {
		return
	}
	adm.Permissions = repositories.NewRolesPermissionsRepo().
		GetPermissionsByRoles(RolesIDUint64...)
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
	return admServ.HasPermission(adm, NewPermissionsServiceByOrm(admServ.repo.GetOrm()).
		GeneratePermissionAuthIdentify(r.Method, r.Path))
}

func (admServ AdminService) HasPermission(adm model.Admin, permIdent string) bool {
	if adm.IsRootRole() {
		return true
	}
	admServ.RefreshPermissions(&adm, false, false)
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
		admServ.RefreshPermissions(&v, false, false)
		_adm = append(_adm, v.ShowMap())
	}
	return _adm
}

// 验证参数 并返回到一个新的adm model
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
			adm.TokenStatus = global.IsNo
		}
		adm.SetRoleID(aValidator.RolesID)
	}

	//这里默认的系统超级管理员 不能被修改角色权限
	if adm.IsRootAccount() {
		if !adm.IsRootRole() {
			return adm, sErr.New("系统默认管理账户权限不能被修改")
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
	ID          uint64
	OrganizerID uint64 `label:"承办单位"`
	Username    string `validate:"required" label:"账户名"`
	Phone       string `validate:"required" label:"手机号码"`
	QQ          string `label:"QQ号" validate:"max=20"`
	Status      uint8  `validate:"oneof=0 1" label:"状态"`
	RealName    string `validate:"required" label:"真实姓名"`
	Avatar      string `label:"头像"`
	Password    string `label:"密码"`
	Sex         uint8  `validate:"oneof=0 1 2" label:"性别"`
	Desc        string `label:"简介"`
	RolesID     []string
	Self        string //用于请求的时候在控制器区分一下是否是编辑当前用户的资料
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
