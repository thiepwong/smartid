package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/services"
)

type AccountController struct {
	Ctx     iris.Context
	Service services.AccountService
}

func (c *AccountController) Get() (results string) {
	results = c.Service.Get()
	return results
}

func AccountHanlder(app *mvc.Application) {
	movieService := services.RegSignupService()
	app.Register(movieService)

	app.Handle(new(AccountController))
}
