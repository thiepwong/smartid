package controllers

import (
	"bytes"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/services"
	"gopkg.in/mgo.v2"
)

type AccountController struct {
	Ctx     iris.Context
	Service services.AccountService
}

// Get func to get all
func (c *AccountController) Get() (results string) {
	//	results = c.Service.Get()
	r := c.Ctx.URLParam("name")
	d := c.Service.Get()
	var msg bytes.Buffer
	fmt.Fprintf(&msg, "Da nhan duoc du lieu %s \r\n va phan hoi la: %s ", r, d)
	return msg.String()
}

func AccountHanlder(app *mvc.Application) {
	dialInfo, err := mgo.ParseURL("mongodb://171.244.49.164:2688/ucenter")

	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	//	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("test")

	myService := services.NewAccountService(db, session)
	app.Register(myService)

	app.Handle(new(AccountController))
}
