package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/controllers"
)

func RegisterRoute(app *iris.Application) {
	mvc.Configure(app.Party("/account"), controllers.AccountHanlder)

}
