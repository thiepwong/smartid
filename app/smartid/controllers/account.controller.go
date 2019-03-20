package controllers

import (
	"errors"
	"fmt"
	"log"

	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/services"
)

//AccountController type
type AccountController struct {
	Ctx            iris.Context
	AccountService services.AccountService
}

//BeforeActivation fuc
func (c *AccountController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/register", "PostSignup")
	b.Handle("POST", "/signin", "PostSignin")
	b.Handle("POST", "/test", "PostTest")
}

// PostSignup method
func (c *AccountController) PostSignup() (results *models.AccountModel, err error) {
	var _signupData = models.SignupModel{}
	er := c.Ctx.ReadJSON(&_signupData)
	if er != nil {
		logger.LogErr.Println(er)
		return
	}

	if _signupData.Username == "" || _signupData.Firstname == "" || _signupData.Lastname == "" || _signupData.Password == "" {
		logger.LogDebug.Println("Signup infomation is invalid!")
		return nil, errors.New("Loi roi")
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

	return acc, err
}

//PostSignin func
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

//PostTest func
func (c *AccountController) PostTest() string {

	var profile models.SignupModel

	a := c.Ctx.URLParam("name")
	profile.Firstname = a
	z := c.AccountService.Test(&profile)
	return z
}

//AccountHandler func
func AccountHandler(app *mvc.Application) {

	app.Handle(new(AccountController))
}
