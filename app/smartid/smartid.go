package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	"github.com/thiepwong/smartid/app/smartid/routes"
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

	app := iris.Default()
	routes.RegisterRoute(app)

	er := app.Run(iris.Addr(*config.host + ":" + strconv.Itoa(*config.port)))
	if er != nil {
		fmt.Println("Server not started!")
	}

}

func loadConfig() (cfg Config, err *error) {

	cfg.system = flag.Int("s", 6, "System code as an integer (1-9), 3, 4, 5 etc. default is 6, is testnet")
	cfg.node = flag.Int("n", 0, "Node code of system, from 0x00-0xFF (001-255), default is 000")
	cfg.host = flag.String("h", "localhost", "Domain or ip of host, default is localhost")
	cfg.port = flag.Int("p", 80, "Port of service, default is 80")
	cfg.cfgpath = flag.String("c", "conf/conf.ini", "Path of program's configuration path, default is conf/conf.ini")
	flag.Parse()
	return cfg, err
}
