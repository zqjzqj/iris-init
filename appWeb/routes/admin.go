package routes

import (
	"big_data_new/appWeb/controller/admin"
	"big_data_new/appWeb/middleware/adminMiddleware"
	"big_data_new/config"
	"big_data_new/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"net/http"
)

func RegisterRoutes(app *iris.Application) {
	tmpl := iris.Django("./views/admin", ".html")
	if !config.EnvIsPro() {
		tmpl = tmpl.Reload(true)
	}
	party := app.Party("/") //.Subdomain("admin")
	sess := sessions.New(sessions.Config{
		Cookie: "big_data_new",
	})
	party.Use(sess.Handler())
	party.RegisterView(tmpl)
	// 注册静态文件
	party.HandleDir("/static", "./static")
	party.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		if !ctx.IsAjax() /*&& !strings.HasPrefix(ctx.Request().URL.Path, "/api/")*/ {
			ctx.Redirect("/err?Msg=页面未找到", http.StatusFound)
		}
	})

	//重要！！！ admin所有需要设置权限的实际路由需要到控制器里去实际注册 并SetName 不然无法获取并生成到路由的权限--会被默认为不需要权限
	//SetName格式 SetName("目录@菜单:按钮") 这里暂时只支持二级菜单
	//b.Handle(http.MethodGet, "list", "GetList").SetName("用户组@用户列表")
	//b.Handle(http.MethodPost, "delete", "PostDelete").SetName("用户组@用户列表:删除")
	//子按钮 子页面按钮点击跳转新的页面
	//b.Handle(http.MethodPost, "delete", "PostDelete").SetName("用户组@用户列表:子页面:子页面")

	//设置排序的方式
	// 默认会按照 application.Party编写的顺序来生成路由菜单的顺序
	// 设置排序方式 b.Handle(http.MethodGet, "list", "GetList").SetName("用户组.Sort{99}@用户列表.Sort{99}")
	// 设置排序方式 b.Handle(http.MethodGet, "list", "GetList").SetName("用户组@用户列表.Sort{99}")
	// 设置排序方式 b.Handle(http.MethodGet, "list", "GetList").SetName("用户组.Sort{99}@用户列表")
	// 目录排序要写在 第一个目录路由的里面 如 admin和role都是用户管理 则用户管理目录排序要写在admin内
	mvc.Configure(party, func(application *mvc.Application) {
		application.Register(adminMiddleware.RegisterAdmin).Handle(&admin.LoginController{})

		application.Party("/", adminMiddleware.AdminLogin, adminMiddleware.AdminPermission).
			Register(adminMiddleware.RegisterAdmin).Handle(&admin.SiteController{})

		application.Party("/settings", adminMiddleware.AdminLogin).Handle(&admin.SettingsController{})

		application.Party("/upload", adminMiddleware.AdminLogin).Handle(&admin.UploadController{})

		application.Party("/areas", adminMiddleware.AdminLogin, adminMiddleware.AdminPermission).Handle(&admin.AreaController{})

		application.Party("/admin", adminMiddleware.AdminLogin, adminMiddleware.AdminPermission).
			Register(adminMiddleware.RegisterAdmin).Handle(&admin.AdminController{})

		application.Party("/roles", adminMiddleware.AdminLogin, adminMiddleware.AdminPermission).
			Register(adminMiddleware.RegisterAdmin).Handle(&admin.RolesController{})
	})

	//刷新一下权限表
	services.NewPermissionsService().GenerateAdminPermissionsByRoutes(app)

}
