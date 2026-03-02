package admin

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-init/appWeb"
	"iris-init/model"
	"iris-init/services"
	"net/http"
)

type RolesController struct {
	Admin model.Admin
}

func (rolesCtrl RolesController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList").SetName("账号管理@角色管理")
	b.Handle(http.MethodGet, "item", "GetItem").SetName("账号管理@角色管理:查看角色")
	b.Handle(http.MethodPost, "edit", "PostEdit").SetName("账号管理@角色管理:编辑角色")
	b.Handle(http.MethodPost, "delete", "PostDelete").SetName("账号管理@角色管理:删除角色")

}

func (rolesCtrl RolesController) GetList(ctx iris.Context) mvc.Result {
	roleServ := services.NewRolesService()
	return appWeb.ResponseDataViewForm("roles/list.html", appWeb.DataView{
		Data: map[string]interface{}{
			"List": roleServ.ShowMapList(roleServ.List(ctx)),
		},
	}, ctx)
}

func (rolesCtrl RolesController) GetItem(ctx iris.Context) mvc.Result {
	rolesServ := services.NewRolesService()
	role := rolesServ.GetItem(ctx)
	rolesServ.RefreshPermission(&role, true)
	permsTreeJson, _ := json.Marshal(services.NewPermissionsService().GetPermTree(role.PermIdents...))
	return appWeb.ResponseDataViewForm("roles/item.html", appWeb.DataView{
		PageJs: []string{"roles/item.js"},
		Data: map[string]interface{}{
			"Role":          role,
			"PermsTreeJson": string(permsTreeJson),
		},
	}, ctx)
}

func (rolesCtrl RolesController) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	role, err := services.NewRolesService().SaveByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", role.ShowMap())
}

func (rolesCtrl RolesController) PostDelete(ctx iris.Context) appWeb.ResponseFormat {
	err := services.NewRolesService().DeleteByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", nil)
}
