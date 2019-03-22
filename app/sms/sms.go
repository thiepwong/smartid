package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/kataras/iris"
	"github.com/thiepwong/smartid/app/sms/routes"
	"github.com/thiepwong/smartid/pkg/config"
)

func main() {

	conf, es := config.LoadConfig()

	if es != nil {
		logger.LogErr.Println(es)
		os.Exit(0)
	}

	fmt.Println(&conf.Environment.System, *conf.Environment.Node, *conf.Environment.Host, *conf.Environment.Port, *conf.Environment.Cfgpath)

	app := iris.New()

	fmt.Println(conf.Database.Redis.Host)
	//	ExampleNewClient()
	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Next()
	}
	app.Logger().SetLevel("debug")
	routes.RegisterRoute(app, crs, &conf)
	fmt.Println("Da chay tai: ", *conf.Environment.Port)

	er := app.Run(iris.Addr(*conf.Environment.Host+":"+strconv.Itoa(*conf.Environment.Port)), iris.WithoutPathCorrectionRedirection)
	if er != nil {
		fmt.Println("Server not started!")
	}

}
