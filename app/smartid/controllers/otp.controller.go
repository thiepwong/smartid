package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/services"
	"github.com/thiepwong/smartid/pkg/logger"
)

type OtpController struct {
	Ctx        iris.Context
	OtpService services.OtpService
	Result     MvcResult
}

//BeforeActivation fuc
func (c *OtpController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/send-otp", "PostSendOTP")
}

func (c *OtpController) PostSendOTP() MvcResult {
	_data := models.NewOTP{}
	er := c.Ctx.ReadJSON(&_data)
	if er != nil {
		logger.LogErr.Println(er)
		c.Result.GenerateResult(500, "Read input data error!", nil)
		return c.Result
	}
	c.Result.GenerateResult(0, "", c.OtpService.SendOTP(_data.Mobile, _data.Ttl))
	return c.Result
}
