package controllers

import (
	"bytes"
	"fmt"
	"log"

	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/smartid/app/smartid/datasources"
	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/services"
	"gopkg.in/mgo.v2"
)

type AccountController struct {
	Ctx     iris.Context
	Service services.AccountService
	Session *mgo.Session
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

type MvcResult struct {
	Version string
	Result  string
}

//BeforeActivation
// Register paths of controllers
func (c *AccountController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/register", "PostSignup")
}

func (c *AccountController) PostSignup() (results models.SignupModel) {
	var _signupData = models.SignupModel{}
	err := c.Ctx.ReadJSON(&_signupData)
	if err != nil {
		log.Fatal()
		return
	}

	e := c.Session.DB("test").C("account").Insert(MvcResult{Result: "Ok", Version: "2.445"})
	if e != nil {
		logger.LogErr.Println(err.Error())
	}
	//	res, err := json.Marshal(_signupData)
	//	d := c.Service.Signup(c.Session, &_signupData)
	return models.SignupModel{}
}

func AccountHanlder(app *mvc.Application) {
	//	dialInfo, err := mgo.ParseURL("mongodb://171.244.49.164:2688/ucenter")

	//	session, err := mgo.DialWithInfo(dialInfo)

	///	s := session.Copy()
	//	if nil != err {
	//	panic(err)
	//	}
	//	defer session.Close()
	//	session.SetMode(mgo.Monotonic, true)
	//	db := session.DB("test")

	s := datasources.GetSession()
	myService := services.NewAccountService(s)
	app.Register(myService)
	app.Handle(new(AccountController))
}
