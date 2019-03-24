package main

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	"github.com/thiepwong/smartid/app/smartid/routes"
	"github.com/thiepwong/smartid/pkg/config"
)

func main() {

	conf, es := config.LoadConfig()

	if es != nil {

	}

	app := iris.New()

	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
		ctx.Next()
	}
	app.Logger().SetLevel("debug")
	routes.RegisterRoute(app, crs, &conf)

	er := app.Run(iris.Addr(*conf.Environment.Host+":"+strconv.Itoa(*conf.Environment.Port)), iris.WithoutPathCorrectionRedirection)
	if er != nil {
		fmt.Println("Server not started!")
	}

}
