package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/sms/services"
)

type SmsController struct {
	Ctx     iris.Context
	Service services.SmsService
	Result  MvcResult
}

//BeforeActivation fuc
func (c *SmsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/login", "PostLogin")
	b.Handle("POST", "/send-sms", "PostSend")
}

func (c *SmsController) PostLogin() MvcResult {
	c.Result.GenerateResult(0, "", c.Ctx.FormValues())
	return c.Result
}

func (c *SmsController) PostSend() MvcResult {
	return c.Result
}
