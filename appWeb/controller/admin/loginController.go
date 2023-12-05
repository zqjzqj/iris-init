package admin

import (
	"github.com/kataras/iris/v12"
	"iris-init/appWeb"
	"iris-init/model"
	"iris-init/services"
)

type LoginController struct {
	Adm model.Admin
}

func (loginCtrl LoginController) PostLogin(ctx iris.Context) appWeb.ResponseFormat {
	admServ := services.NewAdminService()
	adm, err := admServ.LoginByPwd(ctx.PostValueTrim("Username"), ctx.PostValueTrim("Password"))
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", adm.ShowMapHasToken())
}
