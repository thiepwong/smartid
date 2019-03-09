package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"smart/pkg/logger"
	"strconv"

	"github.com/kataras/iris"
)

type Config struct {
	system  *int
	node    *int
	host    *string
	port    *int
	cfgpath *string
}

func main() {

	config, err := loadConfig()

	if err != nil {

	}
	fmt.Println(&config.system, *config.node, *config.host, *config.port, *config.cfgpath) //AR 3700

	app := iris.New()
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
			"name":    &config.system,
			"age":     *config.node})
	})

	fmt.Println("======= Spin Gateway Severvice ========")

	app.Post("/test", func(ctx iris.Context) {
		fmt.Printf("Bien so: %s, Id cam: %s, Thoi gian chup: %s", ctx.FormValue("car_plate"), ctx.FormValue("camera_id"), ctx.FormValue("start_time"))
		ctx.JSON(iris.Map{
			"data": ctx.FormValue("car_plate"),
			"id":   ctx.FormValue("camera_id:")})
	})

	logger.LogInfo.Printf("%d, %s", 545, "Da khoi tao thanh cong")
	//	Log("Da chay thanh cong roi")
	er := app.Run(iris.Addr(*config.host + ":" + strconv.Itoa(*config.port)))
	if er != nil {
		fmt.Println("Server not started!")
	}

}

func loadConfig() (cfg Config, err *error) {
	// var system *int
	// var node *int
	// var host *string
	// var port *string
	// var confPath *string
	//var	system *int, node *int, host *string, port *int, confPath *string, err *error
	cfg.system = flag.Int("s", 6, "System code as an integer (1-9), 3, 4, 5 etc. default is 6, is testnet")
	cfg.node = flag.Int("n", 0, "Node code of system, from 0x00-0xFF (001-255), default is 000")
	cfg.host = flag.String("h", "localhost", "Domain or ip of host, default is localhost")
	cfg.port = flag.Int("p", 80, "Port of service, default is 80")
	cfg.cfgpath = flag.String("c", "conf/conf.ini", "Path of program's configuration path, default is conf/conf.ini")
	flag.Parse()
	return cfg, err
}

func Log(content string) {
	f, err := os.OpenFile("debug.sml", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(content)
}
