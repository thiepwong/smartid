package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/controllers"
	"github.com/thiepwong/smartid/app/smartid/datasources"
	"github.com/thiepwong/smartid/app/smartid/repositories"
	"github.com/thiepwong/smartid/app/smartid/services"
)

func RegisterRoute(app *iris.Application) {

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

	account := mvc.New(app.Party("/account").AllowMethods(iris.MethodPost, iris.MethodOptions))
	//	account := mvc.New(app.Party("/account").AllowMethods(iris.MethodOptions))
	account.Register(accountService)
	account.Handle(new(controllers.AccountController))

	auth := mvc.New(app.Party("/auth"))
	auth.Handle(new(controllers.AuthController))

}
