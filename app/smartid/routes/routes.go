package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/controllers"
	"github.com/thiepwong/smartid/app/smartid/datasources"
	"github.com/thiepwong/smartid/app/smartid/repositories"
	"github.com/thiepwong/smartid/app/smartid/services"
)

type SetHeader func(iris.Context)

func RegisterRoute(app *iris.Application, cors context.Handler) {

	db, err := datasources.GetMongoDb()
	if err != nil {
		fmt.Println("Loi ket noi co so du lieu ", err)
		log.Fatal(err)
		os.Exit(2)
	}
	_c, _e := db.C("accounts").Count()
	if _e != nil {
	}

	fmt.Printf("ten co so du lieu: %d   ", _c)
	accountRepository := repositories.NewAccountRepositoryContext(db, "accounts")
	accountService := services.NewAccountService(accountRepository)
	mvcResult := controllers.NewMvcResult(nil)
	account := mvc.New(app.Party("/account", cors).AllowMethods(iris.MethodOptions))
	account.Register(accountService, mvcResult)
	account.Handle(new(controllers.AccountController))

	auth := mvc.New(app.Party("/auth"))
	auth.Handle(new(controllers.AuthController))

}
