package admin

import (
	"github.com/kataras/iris/v12"
	"iris-init/appWeb"
	"iris-init/services"
)

type AreaController struct {
}

func (area AreaController) GetList(ctx iris.Context) appWeb.ResponseFormat {
	return appWeb.NewSuccessResponse("", services.NewAreaService().GetListByPID(uint(ctx.URLParamUint64("PID"))))
}
