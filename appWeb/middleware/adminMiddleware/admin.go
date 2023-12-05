package adminMiddleware

import (
	"big_data_new/appWeb"
	"big_data_new/global"
	"big_data_new/model"
	"big_data_new/repositories"
	"big_data_new/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"net/http"
)

func RegisterAdminAndAuth(ctx iris.Context) model.Admin {
	adm := RegisterAdmin(ctx)
	if adm.ID == 0 {
		NotLoginHandle(ctx)
		return adm
	}
	return adm
}

func RegisterAdmin(ctx iris.Context) model.Admin {
	r := ctx.Values().Get("adm")
	if r != nil {
		adm, ok := r.(model.Admin)
		if ok {
			return adm
		}
	}
	token := global.GetReqToken(ctx)
	if token == "" {
		return model.Admin{}
	}
	admRepo := repositories.NewAdminRepo()
	adm := admRepo.GetByToken(token)
	if !adm.TokenValid() {
		return model.Admin{}
	}
	ctx.Values().Set("adm", adm)
	return adm
}

func AdminLogin(ctx iris.Context) {
	adm := RegisterAdmin(ctx)
	if adm.ID == 0 {
		//未找到 数据
		NotLoginHandle(ctx)
		return
	}
	ctx.Next()
}

func AdminPermission(ctx iris.Context) {
	adm := RegisterAdmin(ctx)
	if adm.ID == 0 {
		//未找到 数据
		NotLoginHandle(ctx)
		return
	}
	//不是超级管理员
	if !adm.IsRootRole() {
		permServ := services.NewPermissionsService()
		ident := permServ.GeneratePermissionAuthIdentify(ctx.Method(), ctx.Path())
		//这里判断一下此路径是否需要判断权限 不在权限表里 默认通过
		if permServ.IdentifyExists(ident) {
			admServ := services.NewAdminService()
			if !admServ.HasPermission(adm, ident) {
				NotAuthHandle(ctx)
				return
			}
		}
	}
	ctx.Next()
	return
}

func CheckPermissionByRouteFunc(r *router.Route) func(ctx iris.Context) {
	return func(c iris.Context) {
		RegisterAdmin(c)
		if !services.NewAdminService().CheckPermissionByRoute(r, c) {
			_ = c.StopWithJSON(http.StatusOK, appWeb.NewNotAuthResponse("", nil))
			return
		}
		c.Next()
	}
}

func NotAuthHandle(ctx iris.Context) {
	_ = ctx.StopWithJSON(http.StatusOK, appWeb.NewNotAuthResponse("", nil))
	return
}

func NotLoginHandle(ctx iris.Context) {
	_ = ctx.StopWithJSON(http.StatusOK, appWeb.NewNotLoginResponse("", nil))
}
