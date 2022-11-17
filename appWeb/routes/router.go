package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"jd-fxl/appWeb/controller"
	"jd-fxl/appWeb/middleware"
	"jd-fxl/config"
	"jd-fxl/services"
)

func RegisterRoutes(app *iris.Application) {
	tmpl := iris.Django("./views", ".html")
	if !config.EnvIsPro() {
		tmpl = tmpl.Reload(true)
	}
	app.RegisterView(tmpl)
	app.HandleDir("/static", "./static")
	//重要！！！ admin所有需要设置权限的实际路由需要到控制器里去实际注册 并SetName 不然无法获取并生成到路由的权限--会被默认为不需要权限
	//SetName格式 SetName("目录@菜单:按钮") 这里暂时只支持二级菜单
	//b.Handle(http.MethodGet, "list", "GetList").SetName("用户组@用户列表")
	//b.Handle(http.MethodPost, "delete", "PostDelete").SetName("用户组@用户列表:删除")

	//子按钮 子页面按钮点击跳转新的页面
	//b.Handle(http.MethodPost, "delete", "PostDelete").SetName("用户组@用户列表:子页面:子页面")
	mvc.Configure(app.Party("/"), func(application *mvc.Application) {
		application.Register(middleware.RegisterAdmin).Handle(&controller.LoginController{})

		application.Party("/", middleware.AdminLogin, middleware.AdminPermission).
			Register(middleware.RegisterAdmin).Handle(&controller.SiteController{})

		application.Party("/areas", middleware.AdminLogin, middleware.AdminPermission).Handle(&controller.AreaController{})

		application.Party("/admin", middleware.AdminLogin, middleware.AdminPermission).
			Register(middleware.RegisterAdmin).Handle(&controller.AdminController{})

		application.Party("/roles", middleware.AdminLogin, middleware.AdminPermission).
			Register(middleware.RegisterAdmin).Handle(&controller.RolesController{})

	})

	//刷新一下权限表
	services.NewPermissionService().GenerateAdminPermissionsByRoutes(app)

}
