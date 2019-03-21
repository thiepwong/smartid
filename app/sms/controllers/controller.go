package controllers

import "time"

type MvcResult interface {
	GenerateResult(int, string, interface{}) (result *mvcResult)
}

type mvcResult struct {
	System      string
	Version     string
	RequestTime int64
	Code        int
	Message     string
	Data        interface{}
}

func NewMvcResult(result *interface{}) MvcResult {
	return &mvcResult{Data: result, Code: 100, Version: "1.0", System: "SmartID"}

}

func (c *mvcResult) GenerateResult(code int, msg string, d interface{}) (result *mvcResult) {

	if code == 0 {
		code = 200
	}

	if msg == "" {
		msg = "Successful"
	}

	c.RequestTime = time.Now().Unix()
	c.Code = code
	c.Message = msg
	c.Data = d
	return c
}
