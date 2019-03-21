package routes

import (
	"github.com/thiepwong/smartid/app/sms/controllers"
	"github.com/thiepwong/smartid/app/sms/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

type SetHeader func(iris.Context)

func RegisterRoute(app *iris.Application, cors context.Handler) {

	// db, err := datasources.GetMongoDb()
	// if err != nil {
	// 	fmt.Println("Loi ket noi co so du lieu ", err)
	// 	log.Fatal(err)
	// 	os.Exit(2)
	// }
	// _c, _e := db.C("accounts").Count()
	// if _e != nil {
	// }

	// fmt.Printf("ten co so du lieu: %d   ", _c)
	// accountRepository := repositories.NewAccountRepositoryContext(db, "accounts")
	smsService := services.NewSmsService(nil)
	mvcResult := controllers.NewMvcResult(nil)
	account := mvc.New(app.Party("/sms", cors).AllowMethods(iris.MethodOptions))
	account.Register(smsService, mvcResult)
	account.Handle(new(controllers.SmsController))

}
