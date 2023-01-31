package global

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

const (
	ReqTokenName       = "Token"
	ReqTokenHeaderName = "X-Token"
)

func IsApiReq(ctx iris.Context) bool {
	return ctx.URLParamBoolDefault("API", false)
}

func DelCtxUrlParams(ctx iris.Context, keys ...string) {
	req := ctx.Request()
	query := req.URL.Query()
	for _, k := range keys {
		query.Del(k)
	}
	req.URL.RawQuery = query.Encode()
	ctx.ResetRequest(req)
	ctx.ResetQuery()
}

func SetCtxUrlParams(ctx iris.Context, key, val string) {
	req := ctx.Request()
	query := req.URL.Query()
	query.Set(key, val)
	req.URL.RawQuery = query.Encode()
	ctx.ResetRequest(req)
	ctx.ResetQuery()
}

func SetCtxUrlParamsByMap(ctx iris.Context, data map[string]string) {
	req := ctx.Request()
	query := req.URL.Query()
	for k, v := range data {
		query.Set(k, v)
	}
	req.URL.RawQuery = query.Encode()
	ctx.ResetRequest(req)
	ctx.ResetQuery()
}

func GetReqToken(ctx iris.Context) string {
	token := ctx.GetHeader(ReqTokenHeaderName)
	if token == "" {
		token = ctx.URLParamTrim(ReqTokenName)
	}
	if token == "" {
		token = ctx.PostValue(ReqTokenName)
	}
	if token == "" {
		//从session获取
		if sess := sessions.Get(ctx); sess != nil {
			t := sess.Get(ReqTokenName)
			var ok bool
			token, ok = t.(string)
			if !ok || token == "" {
				return ""
			}
		}
	}
	return token
}

func SessionTryDestroy(ctx iris.Context) {
	sess := sessions.Get(ctx)
	if sess != nil {
		sess.Destroy()
	}
}
