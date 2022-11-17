package controller

import (
	"github.com/kataras/iris/v12"
	"jd-fxl/appWeb"
	"jd-fxl/services"
)

type AreaController struct {
}

func (area AreaController) GetList(ctx iris.Context) appWeb.ResponseFormat {
	return appWeb.NewSuccessResponse("", services.NewAreaService().GetListByPID(uint(ctx.URLParamUint64("PID"))))
}
