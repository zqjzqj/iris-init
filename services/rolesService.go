package services

import (
	"github.com/kataras/iris/v12"
	"iris-init/global"
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

func (roleServ RolesService) List(ctx iris.Context) []map[string]interface{} {
	where := repoInterface.RolesSearchWhere{}
	_ = ctx.ReadQuery(&where)
	where.SelectParams = repoComm.SelectFrom{
		OrderBy: []repoComm.OrderByParams{{
			Column: "ID",
			Desc:   true,
		}},
	}
	return roleServ.ShowMapList(roleServ.repo.GetList(where))
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
	err = roleServ.repo.Save(&role)
	return role, err
}

func (roleServ RolesService) DeleteByCtx(ctx iris.Context) error {
	_, err := roleServ.repo.Delete("id", ctx.PostValueInt64Default("ID", 0))
	return err
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
	return role, nil
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
