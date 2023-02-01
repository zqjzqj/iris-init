package admin

import (
	"github.com/kataras/iris/v12"
	"iris-init/appWeb"
	"iris-init/global"
	"strings"
)

type UploadController struct {
}

// 本地上传方法 配合封装好的js使用 注意接收参数File与返回json格式
func (upload UploadController) Post(ctx iris.Context) appWeb.ResponseFormat {
	filename, err := global.UploadLocalByCtx(ctx, "File", "./static/uploads/", "")
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", map[string]interface{}{
		"Url": strings.TrimLeft(filename, "."),
	})
}
