package admin

import (
	"big_data_new/appWeb"
	"big_data_new/model"
	"big_data_new/services"
	"github.com/kataras/iris/v12"
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
