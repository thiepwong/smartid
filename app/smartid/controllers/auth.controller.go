package controllers

import (
	"encoding/json"
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/services"
)

//AuthController struct
type AuthController struct {
	Ctx         iris.Context
	AuthService services.AuthService
}

//BeforeActivation function
func (c *AuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/signin", "PostSignin")
}

//PostSignin func
func (c *AuthController) PostSignin() string {
	var _signinData = models.SigninModel{}
	err := c.Ctx.ReadJSON(&_signinData)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	d, e := json.Marshal(_signinData)
	if e != nil {
		return e.Error()
	}

	return string(d)
}
