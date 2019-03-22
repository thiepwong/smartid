package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/sms/controllers"
	"github.com/thiepwong/smartid/app/sms/datasource"
	"github.com/thiepwong/smartid/app/sms/repositories"
	"github.com/thiepwong/smartid/app/sms/services"
	"github.com/thiepwong/smartid/pkg/config"
)

type SetHeader func(iris.Context)

func RegisterRoute(app *iris.Application, cors context.Handler, config *config.Config) {

	red := datasource.GetRedisDb(config.Database.Redis)
	smsRepository := repositories.NewSmsRepository(red)
	smsService := services.NewSmsService(smsRepository)
	mvcResult := controllers.NewMvcResult(nil)
	account := mvc.New(app.Party("/sms", cors).AllowMethods(iris.MethodOptions))
	account.Register(smsService, mvcResult)
	account.Handle(new(controllers.SmsController))

}
