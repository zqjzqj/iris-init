package admin

import (
	"big_data_new/appWeb"
	"big_data_new/global"
	"big_data_new/ueditor"
	"github.com/kataras/iris/v12"
	"strings"
)

type UploadController struct {
}

// 本地上传方法 配合封装好的js使用 注意接收参数File与返回json格式
func (uploadCtrl UploadController) Post(ctx iris.Context) appWeb.ResponseFormat {
	//处理一下文件名 防止重复了
	filename, err := global.UploadLocalByCtx(ctx, "File", "./static/uploads/", "")
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", map[string]interface{}{
		"Url": strings.TrimLeft(filename, "."),
	})
}

func (uploadCtrl UploadController) AnyUeditor(ctx iris.Context) any {
	uedService, _ := ueditor.NewService(nil, nil, "./", "")
	switch ctx.URLParamTrim("action") {
	case "config":
		return uedService.Config()
	case "uploadimage":
		// 上传图片
		r, _ := uedService.Uploadimage(ctx.Request())
		return r
	case "uploadscrawl":
		// 上传涂鸦
		r, _ := uedService.UploadScrawl(ctx.Request())
		return r
	case "uploadvideo":
		// 上传视频
		r, _ := uedService.UploadVideo(ctx.Request())
		return r
	case "uploadfile":
		// 上传附件
		r, _ := uedService.UploadFile(ctx.Request())
		return r
	case "listfile":
		// 查询上传的文件列表
		listFileItem := ueditor.ListFileItem{}
		uedService.Listfile(&listFileItem, 0, 10)
		return listFileItem
	case "listimage":
		listFileItem := ueditor.ListFileItem{}
		uedService.ListImage(&listFileItem, 0, 10)
		// 查询上传的图片列表
		return listFileItem
	}
	return appWeb.NewSuccessResponse("", nil)
}
