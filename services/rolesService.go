package services

import (
	"big_data_new/global"
	"big_data_new/logs"
	"big_data_new/model"
	"big_data_new/repositories"
	"big_data_new/repositories/repoComm"
	"big_data_new/repositories/repoInterface"
	"big_data_new/sErr"
	"github.com/kataras/iris/v12"
)

func NewRolesService() RolesService {
	return RolesService{repo: repositories.NewRolesRepo()}
}

func NewRolesServiceByOrm(orm any) RolesService {
	r := RolesService{repo: repositories.NewRolesRepo()}
	r.repo.SetOrm(orm)
	return r
}

func NewRolesServiceByRepo(repo repoInterface.RolesRepo) RolesService {
	return RolesService{repo: repo}
}

type RolesService struct {
	repo repoInterface.RolesRepo
}

func (rolesServ RolesService) ListPage(ctx iris.Context) ([]model.Roles, *global.Pager) {
	where := repoInterface.RolesSearchWhere{}
	_ = ctx.ReadQuery(&where)
	pager := global.NewPager(ctx)
	if pager.Size < 0 {
		return rolesServ.repo.GetList(where), nil
	}
	pager.SetTotal(rolesServ.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []model.Roles{}, pager
	}
	where.SelectParams = repoComm.SelectFrom{
		Offset:  pager.Offset,
		Limit:   pager.Size,
		RetSize: pager.Size,
		OrderBy: []repoComm.OrderByParams{
			{
				Column: "ID",
				Desc:   true,
			},
		},
	}
	return rolesServ.repo.GetList(where), pager
}

func (rolesServ RolesService) ListAvailable(_select ...string) []model.Roles {
	if len(_select) == 0 {
		_select = nil
	}
	return rolesServ.repo.GetList(repoInterface.RolesSearchWhere{
		SelectParams: repoComm.SelectFrom{
			Select: _select,
			OrderBy: []repoComm.OrderByParams{{
				Column: "ID",
				Desc:   true,
			}},
		},
	})
}

func (rolesServ RolesService) ListByWhere(where repoInterface.RolesSearchWhere) []model.Roles {
	return rolesServ.repo.GetList(where)
}

func (rolesServ RolesService) TotalCount(where repoInterface.RolesSearchWhere) int64 {
	return rolesServ.repo.GetTotalCount(where)
}

func (roleServ RolesService) GetItem(ctx iris.Context) model.Roles {
	return roleServ.repo.GetByID(ctx.URLParamUint64("ID"))
}

func (roleServ RolesService) RefreshPermission(role *model.Roles, force bool) {
	if !force && role.PermIdents != nil {
		return
	}
	role.PermIdents = repositories.NewRolesPermissionsRepo().GetPermissionsByRoles(role.ID)
}

func (roleServ RolesService) List(ctx iris.Context) []model.Roles {
	where := repoInterface.RolesSearchWhere{}
	_ = ctx.ReadQuery(&where)
	where.SelectParams = repoComm.SelectFrom{
		OrderBy: []repoComm.OrderByParams{{
			Column: "ID",
			Desc:   true,
		}},
	}
	return roleServ.repo.GetList(where)
}

func (roleServ RolesService) ShowMapList(roles []model.Roles) []map[string]interface{} {
	_roles := []map[string]interface{}{}
	for _, v := range roles {
		_roles = append(_roles, v.ShowMap())
	}
	return _roles
}

func (roleServ RolesService) SaveByCtx(ctx iris.Context) (model.Roles, error) {
	roleValidator := RolesValidator{}
	err := ctx.ReadBody(&roleValidator)
	if err != nil {
		return model.Roles{}, err
	}
	return roleServ.SaveByValidator(roleValidator)
}

func (roleServ RolesService) SaveByValidator(roleValidator RolesValidator) (model.Roles, error) {
	role, err := roleServ.GetRoleByValidate(roleValidator)
	if err != nil {
		return role, err
	}
	err = roleServ.Save(&role)
	return role, err
}

func (roleServ RolesService) Save(role *model.Roles, _select ...string) error {
	if len(role.PermIdents) == 0 {
		return roleServ.repo.Save(role)
	}
	rPermRepo := repositories.NewRolesPermissionsRepo()
	return roleServ.repo.Transaction(func() error {
		err := roleServ.repo.Save(role, _select...)
		if err != nil {
			return err
		}
		return rPermRepo.SaveByRole(*role)
	}, rPermRepo)
}

func (roleServ RolesService) Delete(role model.Roles) error {
	if role.ID == 0 {
		return nil
	}
	rPermRepo := repositories.NewRolesPermissionsRepo()
	roleAdmRepo := repositories.NewRolesAdmRepo()
	return roleServ.repo.Transaction(func() error {
		for _, v := range model.GetSysRoles() {
			if v.ID == role.ID {
				return sErr.NewFmt("不允许删除内置角色%s", v.Name)
			}
		}
		role.PermIdents = nil
		_err := rPermRepo.SaveByRole(role)
		if _err != nil {
			return _err
		}
		_, _err = roleAdmRepo.DeleteByRoleID(role.ID)
		if _err != nil {
			return _err
		}
		_, _err = roleServ.repo.DeleteByID(role.ID)
		return nil
	}, rPermRepo, roleAdmRepo)
}

func (roleServ RolesService) DeleteByCtx(ctx iris.Context) error {
	return roleServ.Delete(roleServ.repo.GetByID(uint64(ctx.PostValueInt64Default("ID", 0))))
}

func (roleServ RolesService) GetRoleByValidate(roleValidator RolesValidator) (model.Roles, error) {
	err := roleValidator.Validate()
	if err != nil {
		return model.Roles{}, err
	}
	var role model.Roles
	if roleValidator.ID > 0 {
		role = roleServ.repo.GetByID(roleValidator.ID)
		if role.ID == 0 {
			return role, sErr.New("无效的ID")
		}
	} else {
		role = model.Roles{}
	}
	role.Name = roleValidator.Name
	role.Remark = roleValidator.Remark
	role.PermIdents = roleValidator.Idents
	//这里判断如果是内置角色 则将内置权限压入数组 防止权限丢失
	sysRoles := model.GetSysRoles()
	for _, v := range sysRoles {
		if v.ID == role.ID {
			if role.PermIdents == nil {
				role.PermIdents = make([]string, 0, len(sysRoles)*4)
			}
			//这里在save的时候会自动去过滤role.PermIdents中的重复权限 所以直接append即可
			role.PermIdents = append(role.PermIdents, v.PermIdents...)
		}
	}
	return role, nil
}

// 在cmd/updateSysRoles.go中调用
func (roleServ RolesService) UpdateSysRole() {
	permissionServ := NewPermissionsServiceByOrm(roleServ.repo.GetOrm())
	for _, v := range model.GetSysRoles() {
		role := roleServ.repo.GetByID(v.ID)
		permIdents := make([]string, 0, 5)
		for _, _vPermIdent := range v.PermIdents {
			for _, permIdent := range permissionServ.GetPermParentsByIdent(_vPermIdent) {
				permIdents = append(permIdents, permIdent.Ident)
			}
		}
		v.PermIdents = append(v.PermIdents, permIdents...)
		if role.ID > 0 {
			logs.PrintlnInfo("update role ", role.Name, role.PermIdents)
			if role.PermIdents == nil {
				role.PermIdents = make([]string, 0, len(v.PermIdents))
			}
			role.PermIdents = append(role.PermIdents, v.PermIdents...)
			_ = roleServ.Save(&role)
		} else {
			logs.PrintlnInfo("create role ", v.Name)
			_ = roleServ.Save(&v)
		}
	}
}

func (rolesServ RolesService) GetByWhere(where repoInterface.RolesSearchWhere) model.Roles {
	return rolesServ.repo.GetByWhere(where)
}

func (rolesServ RolesService) ScanByWhere(where repoInterface.RolesSearchWhere, dest any) error {
	return rolesServ.repo.ScanByWhere(where, dest)
}

func (rolesServ RolesService) ScanByOrWhere(dest any, where ...repoInterface.RolesSearchWhere) error {
	return rolesServ.repo.ScanByOrWhere(dest, where...)
}

func (rolesServ RolesService) UpdateByWhere(where repoInterface.RolesSearchWhere, data interface{}) (rowsAffected int64, err error) {
	return rolesServ.repo.UpdateByWhere(where, data)
}
func (rolesServ RolesService) GetByName(name string, _select ...string) []model.Roles {
	return rolesServ.repo.GetByName(name, _select...)
}

func (rolesServ RolesService) DeleteByName(name string) error {
	_, err := rolesServ.repo.DeleteByName(name)
	return err
}

func (rolesServ RolesService) DeleteByID(ID ...uint64) error {
	_, err := rolesServ.repo.DeleteByID(ID...)
	return err
}

func (rolesServ RolesService) Create(roles *[]model.Roles) error {
	return rolesServ.repo.Create(roles)
}

func (rolesServ RolesService) GetByIDLock(ID uint64, _select ...string) (model.Roles, repoComm.ReleaseLock) {
	return rolesServ.repo.GetByIDLock(ID, _select...)
}

func (rolesServ RolesService) GetRolesByID(id ...uint64) []model.Roles {
	return rolesServ.repo.GetRolesByID(id...)
}

type RolesValidator struct {
	ID     uint64
	Name   string `validate:"required,max=100"`
	Remark string `validate:"omitempty,max=241"`
	Idents []string
}

func (roleValidator RolesValidator) Validate() error {
	if len(roleValidator.Idents) == 0 {
		return sErr.New("权限不能为空")
	}
	return global.ValidateV9Struct(roleValidator)
}
