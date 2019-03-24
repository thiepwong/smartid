package routes

import (
	"os"

	"github.com/thiepwong/smartid/pkg/config"
	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/controllers"
	"github.com/thiepwong/smartid/app/smartid/datasources"
	"github.com/thiepwong/smartid/app/smartid/repositories"
	"github.com/thiepwong/smartid/app/smartid/services"
	"github.com/thiepwong/smartid/app/sms/datasource"
)

type SetHeader func(iris.Context)

func RegisterRoute(app *iris.Application, cors context.Handler, config *config.Config) {
	red := datasource.GetRedisDb(config.Database.Redis)
	db, err := datasources.GetMongoDb()
	if err != nil {
		logger.LogErr.Println(err.Error())
		os.Exit(2)
	}

	mvcResult := controllers.NewMvcResult(nil)

	// Register the account controller
	accountRepository := repositories.NewAccountRepositoryContext(db, "accounts")
	accountService := services.NewAccountService(accountRepository)
	account := mvc.New(app.Party("/account", cors).AllowMethods(iris.MethodOptions))
	account.Register(accountService, mvcResult)
	account.Handle(new(controllers.AccountController))

	// Register the OTP controller
	otpRepository := repositories.NewOtpRepository(red)
	otpService := services.NewOtpService(otpRepository)
	otp := mvc.New(app.Party("/validate/otp", cors).AllowMethods(iris.MethodOptions))
	otp.Register(otpService, mvcResult)
	otp.Handle(new(controllers.OtpController))

	//Register the Auth
	auth := mvc.New(app.Party("/auth"))
	auth.Handle(new(controllers.AuthController))

}
