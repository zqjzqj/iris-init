package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-init/appWeb"
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

// 获取数据列表
func (admCtrl AdminController) GetList(ctx iris.Context) appWeb.ResponseFormat {
	admServ := services.NewAdminService()
	adm, page := admServ.ListPage(ctx)
	return appWeb.NewPagerResponse(admServ.ShowMapList(adm), page)
}

// 获取一条详细数据
func (admCtrl AdminController) GetItem(ctx iris.Context) appWeb.ResponseFormat {
	admServ := services.NewAdminService()
	adm := admServ.GetItem(ctx)
	if adm.ID == 0 {
		return appWeb.NewFailResponse("无效的数据", nil)
	}
	admServ.RefreshPermissions(&adm, false, false)
	return appWeb.NewSuccessResponse("", adm.ShowMap())
}

func (admCtrl AdminController) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	admServ := services.NewAdminService()
	adm, err := admServ.SaveByCtx(ctx, admCtrl.Admin.ID)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	admServ.RefreshPermissions(&adm, false, false)
	return appWeb.NewSuccessResponse("", adm.ShowMap())
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
	return appWeb.NewSuccessResponse("", nil)
}
