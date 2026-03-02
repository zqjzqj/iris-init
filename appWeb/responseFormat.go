package appWeb

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"iris-init/global"
	"net/url"
)

const (
	ResponseSuccessCode  = 0
	ResponseFailCode     = 1
	ResponseNotLoginCode = -1
	ResponseNotAuthCode  = -2

	AjaxLocationKey          = "_url"
	AjaxLocationLayerOpenKey = "_open_url"
)

type ResponseFormat struct {
	Code int         `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data"`
}

func NewResponse(code int, msg string, data interface{}) ResponseFormat {
	if msg == "" {
		if code == ResponseSuccessCode {
			msg = "操作成功"
		} else if code == ResponseFailCode {
			msg = "操作失败"
		} else if code == ResponseNotLoginCode {
			msg = "账户未登录"
		} else if code == ResponseNotAuthCode {
			msg = "无权访问"
		}
	}
	return ResponseFormat{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewPagerResponse(data interface{}, pager *global.Pager) ResponseFormat {
	return NewResponse(ResponseSuccessCode, "", map[string]interface{}{
		"Page":    pager,
		"Content": data,
	})
}

func NewSuccessResponse(msg string, data interface{}) ResponseFormat {
	return NewResponse(ResponseSuccessCode, msg, data)
}

func NewFailResponse(msg string, data interface{}) ResponseFormat {
	return NewResponse(ResponseFailCode, msg, data)
}

func NewFailErrResponse(err error, data interface{}) ResponseFormat {
	return NewResponse(ResponseFailCode, err.Error(), data)
}

func NewNotAuthResponse(msg string, data interface{}) ResponseFormat {
	return NewResponse(ResponseNotAuthCode, msg, data)
}

func NewNotLoginResponse(msg string, data interface{}) ResponseFormat {
	return NewResponse(ResponseNotLoginCode, msg, data)
}

func RespDownloadFile(ctx iris.Context, filename string, buf *bytes.Buffer) error {
	//ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("X-Filename", url.QueryEscape(filename))
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	ctx.Header("Content-Length", fmt.Sprintf("%d", buf.Len()))
	if _, err := io.Copy(ctx.ResponseWriter(), bytes.NewReader(buf.Bytes())); err != nil {
		return err
	}
	return nil
}

type AjaxLocationLayerOpenParams struct {
	Content string
	Title   string
	Width   string
	Height  string
	Type    string
}
