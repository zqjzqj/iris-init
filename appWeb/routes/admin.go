package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-init/appWeb/controller/admin"
	"iris-init/appWeb/middleware"
	"iris-init/config"
	"iris-init/services"
	"net/http"
)

func RegisterRoutes(app *iris.Application) {
	tmpl := iris.Django("./views/admin", ".html")
	if !config.EnvIsPro() {
		tmpl = tmpl.Reload(true)
	}
	app.RegisterView(tmpl)
	//重要！！！ admin所有需要设置权限的实际路由需要到控制器里去实际注册 并SetName 不然无法获取并生成到路由的权限--会被默认为不需要权限
	//SetName格式 SetName("目录@菜单:按钮") 这里暂时只支持二级菜单
	//b.Handle(http.MethodGet, "list", "GetList").SetName("用户组@用户列表")
	//b.Handle(http.MethodPost, "delete", "PostDelete").SetName("用户组@用户列表:删除")
	//子按钮 子页面按钮点击跳转新的页面
	//b.Handle(http.MethodPost, "delete", "PostDelete").SetName("用户组@用户列表:子页面:子页面")
	party := app.Party("/") //.Subdomain("admin")
	party.HandleDir("/static", "./static")
	party.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		if !ctx.IsAjax() /*&& !strings.HasPrefix(ctx.Request().URL.Path, "/api/")*/ {
			ctx.Redirect("/err?Msg=页面未找到", http.StatusFound)
		}
	})
	mvc.Configure(party, func(application *mvc.Application) {
		application.Register(middleware.RegisterAdmin).Handle(&admin.LoginController{})

		application.Party("/", middleware.AdminLogin, middleware.AdminPermission).
			Register(middleware.RegisterAdmin).Handle(&admin.SiteController{})

		application.Party("/areas", middleware.AdminLogin, middleware.AdminPermission).Handle(&admin.AreaController{})

		application.Party("/admin", middleware.AdminLogin, middleware.AdminPermission).
			Register(middleware.RegisterAdmin).Handle(&admin.AdminController{})

		application.Party("/roles", middleware.AdminLogin, middleware.AdminPermission).
			Register(middleware.RegisterAdmin).Handle(&admin.RolesController{})

	})

	//刷新一下权限表
	services.NewPermissionService().GenerateAdminPermissionsByRoutes(app)

}
