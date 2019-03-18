package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/services"
)

type AccountController struct {
	Ctx            iris.Context
	AccountService services.AccountService
}

//BeforeActivation
// Register paths of controllers
func (c *AccountController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/register", "PostSignup")
	b.Handle("POST", "/signin", "PostSignin")
	b.Handle("POST", "/test", "PostTest")
}

func (c *AccountController) PostSignup() (results string) {
	var _signupData = models.SignupModel{}
	er := c.Ctx.ReadJSON(&_signupData)
	if er != nil {
		log.Fatal()
		return
	}

	var _sign models.SignupModel
	_sign.Firstname = _signupData.Firstname
	_sign.Lastname = _signupData.Lastname
	_sign.Username = _signupData.Username
	_sign.Password = _signupData.Password

	acc, err := c.AccountService.RegisterAccount(&_sign)

	if err != nil {
		logger.LogErr.Println(err)
	}

	res, e := json.Marshal(acc)
	if e != nil {
		logger.LogErr.Println(e)
	}
	return string(res)
}

func (c *AccountController) PostSignin() (results string) {
	var _signinData = models.SigninModel{}
	err := c.Ctx.ReadJSON(&_signinData)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Kiem tra bien C = %x", &c)

	c.AccountService.SigninAccount(_signinData)

	return "Da dang nhap"
}

func (c *AccountController) PostTest() string {

	var profile models.SignupModel

	a := c.Ctx.URLParam("name")
	profile.Firstname = a
	z := c.AccountService.Test(&profile)
	return z
}

// func AccountHanlder(app *mvc.Application) {

// 	//app.Register(accountService)

// 	app.Handle(new(AccountController))
// }
