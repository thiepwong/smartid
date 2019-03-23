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
	kq := c.Service.Login()
	c.Result.GenerateResult(0, "", kq)
	return c.Result
}

func (c *SmsController) PostSend() MvcResult {
	msg := c.Ctx.FormValue("id")
	mobile := c.Ctx.FormValue("mobile")
	k := c.Service.SendSms(mobile, msg)
	c.Result.GenerateResult(0, "", k)
	return c.Result
}
