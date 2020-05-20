package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/ziqiango/goadmin/web/admin/controllers"
)

func Configure(app *iris.Application)  {
	admin := app.Party("/admin")
	{
		indexController := new(controllers.Index)
		admin.Get("", indexController.Index).Name = "admin.index.index"
		admin.Get("/hello", indexController.Hello).Name = "admin.index.hello"
	}
}
