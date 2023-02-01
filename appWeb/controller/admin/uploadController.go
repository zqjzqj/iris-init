package admin

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"io"
	"iris-init/appWeb"
	"iris-init/global"
	"os"
	"strings"
)

type UploadController struct {
}

// 本地上传方法 配合封装好的js使用 注意接收参数File与返回json格式
func (upload UploadController) Post(ctx iris.Context) appWeb.ResponseFormat {
	file, info, err := ctx.FormFile("File")
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	defer func() { _ = file.Close() }()
	//处理一下文件名 防止重复了
	var filename string
	_suffix := strings.Split(info.Filename, ".")
	_suffixLen := len(_suffix)
	if _suffixLen > 1 {
		filename = global.Md5(uuid.New().String()) + "." + _suffix[_suffixLen-1]
	} else {
		filename = global.Md5(uuid.New().String())
	}
	filename = "./static/uploads/" + filename
	saveFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	defer func() { _ = saveFile }()
	_, err = io.Copy(saveFile, file)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", map[string]interface{}{
		"Url": strings.TrimLeft(filename, "."),
	})
}
