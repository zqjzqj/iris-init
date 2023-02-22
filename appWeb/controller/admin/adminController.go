package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-init/appWeb"
	"iris-init/appWeb/resourcePkg"
	"iris-init/model"
	"iris-init/services"
	"net/http"
)

type AdminController struct {
	Admin model.Admin
}

func (admCtrl AdminController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList").SetName("账号管理@账号列表")
	b.Handle(http.MethodGet, "item", "GetItem").SetName("账号管理@账号列表:查看账号")
	b.Handle(http.MethodPost, "edit", "PostEdit").SetName("账号管理@账号列表:编辑账号")
	b.Handle(http.MethodPost, "delete", "PostDelete").SetName("账号管理@账号列表:删除账号")
}

func (admCtrl AdminController) GetPerms() appWeb.ResponseFormat {
	services.NewAdminService().RefreshPermissions(&admCtrl.Admin, false, false)
	return appWeb.NewSuccessResponse("", admCtrl.Admin.Permissions)
}

func (admCtrl AdminController) GetSelf(ctx iris.Context) mvc.Result {
	roleServ := services.NewRolesService()
	return appWeb.ResponseDataViewForm("admin/item.html", appWeb.DataView{
		Data: map[string]interface{}{
			"Adm":   admCtrl.Admin.ShowMap(),
			"Roles": roleServ.ShowMapList(roleServ.List(ctx)),
			"Self":  "1",
		},
		ResourcePkg: []appWeb.ResourcePkg{resourcePkg.Ueditor{}},
	}, ctx)
}

func (admCtrl AdminController) PostSelf(ctx iris.Context) appWeb.ResponseFormat {
	_adm, err := services.NewAdminService().SaveByCtx(ctx, admCtrl.Admin.ID)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", _adm.ShowMap())
}

// 获取数据列表
func (admCtrl AdminController) GetList(ctx iris.Context) mvc.Result {
	admServ := services.NewAdminService()
	adm, page := admServ.ListPage(ctx)
	return appWeb.ResponseDataViewForm("admin/list.html", appWeb.DataView{
		Pager: page,
		Data: map[string]interface{}{
			"List": admServ.ShowMapList(adm),
		},
	}, ctx)
}

// 获取一条详细数据
func (admCtrl AdminController) GetItem(ctx iris.Context) mvc.Result {
	roleServ := services.NewRolesService()
	return appWeb.ResponseDataViewForm("admin/item.html", appWeb.DataView{
		Data: map[string]interface{}{
			"Adm":   services.NewAdminService().GetItem(ctx).ShowMap(),
			"Roles": roleServ.ShowMapList(roleServ.List(ctx)),
		},
		ResourcePkg: []appWeb.ResourcePkg{resourcePkg.Ueditor{}},
	}, ctx)
}

func (admCtrl AdminController) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	_adm, err := services.NewAdminService().SaveByCtx(ctx, 0)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", _adm.ShowMap())
}

func (admCtrl AdminController) PostDelete(ctx iris.Context) appWeb.ResponseFormat {
	if uint64(ctx.PostValueInt64Default("ID", 0)) == admCtrl.Admin.ID {
		return appWeb.NewFailResponse("不能删除当前用户", nil)
	}
	err := services.NewAdminService().DeleteByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", nil)
}

// 注销登录
func (admCtrl AdminController) GetLogout() appWeb.ResponseFormat {
	err := services.NewAdminService().Logout(&admCtrl.Admin)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", map[string]string{
		appWeb.AjaxLocationKey: "/login",
	})
}
