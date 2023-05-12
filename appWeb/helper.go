package appWeb

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-init/global"
	"iris-init/model"
	"iris-init/services"
	"net/url"
	"strings"
)

type ResourcePkg interface {
	GetJs() []string
	GetCss() []string
}

type DataView struct {
	Data        map[string]interface{}
	Logo        string
	FooterLogo  string
	Pager       *global.Pager
	Title       string
	PageCss     []string
	PageJs      []string
	PageUrl     *url.URL
	ShorFooter  bool
	ResourcePkg []ResourcePkg
}

func ResponseView(view mvc.View) mvc.Result {
	return view
}

func ResponseErrView(msg string, ctx iris.Context) mvc.Result {
	if msg == "" {
		msg = "好像出错了呢"
	}
	return ResponseDataView("err/err.html", DataView{
		Data: map[string]interface{}{
			"Msg": msg,
		},
	}, ctx)
}

func ResponseDataView(view string, dataView DataView, ctx iris.Context) mvc.Result {
	if dataView.Data == nil {
		dataView.Data = make(map[string]interface{})
	}
	dataView.Title = services.NewSettingsService().GetWebsiteTitle()
	//引入资源
	if len(dataView.ResourcePkg) > 0 {
		for _, rPkg := range dataView.ResourcePkg {
			dataView.PageCss = append(dataView.PageCss, rPkg.GetCss()...)
			dataView.PageJs = append(dataView.PageJs, rPkg.GetJs()...)
		}
	}
	for k := range dataView.PageCss {
		if !strings.HasPrefix(dataView.PageCss[k], "/static") &&
			!strings.HasPrefix(dataView.PageCss[k], "http") {
			dataView.PageCss[k] = fmt.Sprintf("/static/views/css/%s", dataView.PageCss[k])
		}
	}
	for k := range dataView.PageJs {
		if !strings.HasPrefix(dataView.PageJs[k], "/static") &&
			!strings.HasPrefix(dataView.PageJs[k], "http") {
			dataView.PageJs[k] = fmt.Sprintf("/static/views/js/%s", dataView.PageJs[k])
		}
	}
	//模板注入公共参数
	if ctx != nil {
		dataView.PageUrl = ctx.Request().URL
		r := ctx.Values().Get("adm")
		if r != nil {
			adm, ok := r.(model.Admin)
			if ok {
				dataView.Data["CurrAdm"] = adm.ShowMap()
			}
		}
	}
	return ResponseView(mvc.View{
		Name: view,
		Data: dataView,
	})
}

// 带form提交的视图
func ResponseDataViewForm(view string, dataView DataView, ctx iris.Context) mvc.Result {
	js := []string{
		"formComponents/jquery.form.min.js",
		"/static/layui/lay/modules/layer.js",
		"formComponents/form.callback.js",
		"formComponents/form.alert.js",
		"formComponents/form.common.js",
	}
	if dataView.PageJs == nil {
		dataView.PageJs = js
	} else {
		dataView.PageJs = append(js, dataView.PageJs...)
	}

	return ResponseDataView(view, dataView, ctx)
}
