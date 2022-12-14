package services

import (
	"github.com/kataras/iris/v12"
	"iris-init/global"
	"iris-init/logs"
	"iris-init/model"
	"iris-init/repositories"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
	"iris-init/sErr"
)

func NewRolesService() RolesService {
	return RolesService{repo: repositories.NewRolesRepo()}
}

type RolesService struct {
	repo repoInterface.RolesRepo
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

func (roleServ RolesService) EditByCtx(ctx iris.Context) (model.Roles, error) {
	roleValidator := RolesValidator{}
	err := ctx.ReadBody(&roleValidator)
	if err != nil {
		return model.Roles{}, err
	}
	return roleServ.EditByValidator(roleValidator)
}

func (roleServ RolesService) EditByValidator(roleValidator RolesValidator) (model.Roles, error) {
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
				return sErr.NewFmt("???????????????????????????%s", v.Name)
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
			return role, sErr.New("?????????ID")
		}
	} else {
		role = model.Roles{}
	}
	role.Name = roleValidator.Name
	role.Remark = roleValidator.Remark
	role.PermIdents = roleValidator.Idents
	//????????????????????????????????? ?????????????????????????????? ??????????????????
	sysRoles := model.GetSysRoles()
	for _, v := range sysRoles {
		if v.ID == role.ID {
			if role.PermIdents == nil {
				role.PermIdents = make([]string, 0, len(sysRoles)*4)
			}
			//?????????save???????????????????????????role.PermIdents?????????????????? ????????????append??????
			role.PermIdents = append(role.PermIdents, v.PermIdents...)
		}
	}
	return role, nil
}

// ???cmd/updateSysRoles.go?????????
func (roleServ RolesService) UpdateSysRole() {
	for _, v := range model.GetSysRoles() {
		_v := roleServ.repo.GetByID(v.ID)
		if _v.ID > 0 {
			logs.PrintlnInfo("update role ", _v.Name)
			if _v.PermIdents == nil {
				_v.PermIdents = make([]string, 0, len(v.PermIdents))
			}
			_v.PermIdents = append(_v.PermIdents, v.PermIdents...)
			_ = roleServ.Save(&_v)
		} else {
			logs.PrintlnInfo("create role ", v.Name)
			_ = roleServ.Save(&v)
		}
	}
}

type RolesValidator struct {
	ID     uint64
	Name   string `validate:"required,max=100"`
	Remark string `validate:"omitempty,max=241"`
	Idents []string
}

func (roleValidator RolesValidator) Validate() error {
	if len(roleValidator.Idents) == 0 {
		return sErr.New("??????????????????")
	}
	return global.ValidateV9Struct(roleValidator)
}
