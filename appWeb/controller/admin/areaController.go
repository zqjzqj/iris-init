package admin

import (
	"big_data_new/appWeb"
	"big_data_new/services"
	"github.com/kataras/iris/v12"
)

type AreaController struct {
}

func (areaCtrl AreaController) GetList(ctx iris.Context) appWeb.ResponseFormat {
	return appWeb.NewSuccessResponse("", services.NewAreaService().GetListByPID(uint(ctx.URLParamUint64("PID"))))
}

func (areaCtrl AreaController) GetLayered() appWeb.ResponseFormat {
	return appWeb.NewSuccessResponse("", services.NewAreaService().ListLayered())
}
