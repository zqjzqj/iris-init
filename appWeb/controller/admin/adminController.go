package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"jd-fxl/appWeb"
	"jd-fxl/model"
	"jd-fxl/services"
	"net/http"
)

type AdminController struct {
	Admin model.Admin
}

func (adm AdminController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList").SetName("账号管理@账号列表")
	b.Handle(http.MethodGet, "item", "GetItem").SetName("账号管理@账号列表:查看账号")
	b.Handle(http.MethodPost, "edit", "PostEdit").SetName("账号管理@账号列表:编辑账号")
	b.Handle(http.MethodPost, "delete", "PostDelete").SetName("账号管理@账号列表:删除账号")

}

func (adm AdminController) GetPerms() appWeb.ResponseFormat {
	services.NewAdminService().RefreshPermissions(&adm.Admin, false)
	return appWeb.NewSuccessResponse("", adm.Admin.Permissions)
}

func (adm AdminController) GetSelf(ctx iris.Context) mvc.Result {
	return appWeb.ResponseDataViewForm("admin/item.html", appWeb.DataView{
		Data: map[string]interface{}{
			"Adm":   adm.Admin.ShowMap(),
			"Roles": services.NewRolesService().List(ctx),
			"Self":  "1",
		},
	}, ctx)
}

func (adm AdminController) PostSelf(ctx iris.Context) appWeb.ResponseFormat {
	_adm, err := services.NewAdminService().EditByCtx(ctx, adm.Admin.ID)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", _adm.ShowMap())
}

// 获取数据列表
func (adm AdminController) GetList(ctx iris.Context) mvc.Result {
	list, page := services.NewAdminService().ListPage(ctx)
	return appWeb.ResponseDataViewForm("admin/list.html", appWeb.DataView{
		Pager: page,
		Data: map[string]interface{}{
			"List": list,
		},
	}, ctx)
}

// 获取一条详细数据
func (adm AdminController) GetItem(ctx iris.Context) mvc.Result {
	return appWeb.ResponseDataViewForm("admin/item.html", appWeb.DataView{
		Data: map[string]interface{}{
			"Adm":   services.NewAdminService().GetItem(ctx).ShowMap(),
			"Roles": services.NewRolesService().List(ctx),
		},
	}, ctx)
}

func (adm AdminController) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	_adm, err := services.NewAdminService().EditByCtx(ctx, 0)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", _adm.ShowMap())
}

func (adm AdminController) PostDelete(ctx iris.Context) appWeb.ResponseFormat {
	if uint64(ctx.PostValueInt64Default("ID", 0)) == adm.Admin.ID {
		return appWeb.NewFailResponse("不能删除当前用户", nil)
	}
	err := services.NewAdminService().DeleteByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", nil)
}

// 注销登录
func (adm AdminController) GetLogout() appWeb.ResponseFormat {
	err := services.NewAdminService().Logout(&adm.Admin)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", map[string]string{
		appWeb.AjaxLocationKey: "/login",
	})
}
