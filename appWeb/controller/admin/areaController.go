package admin

import (
	"9xbet_risk/appWeb"
	"9xbet_risk/services"
	"github.com/kataras/iris/v12"
)

type AreaController struct {
}

func (areaCtrl AreaController) GetList(ctx iris.Context) appWeb.ResponseFormat {
	return appWeb.NewSuccessResponse("", services.NewAreaService().GetListByPID(uint(ctx.URLParamUint64("PID"))))
}
