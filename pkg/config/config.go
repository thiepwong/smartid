package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/thiepwong/smartid/pkg/logger"
)

type Config struct {
	Environment *CliConfig
	Vendor      *Vendor
	Database    *Database
}

type Vendor struct {
	Url         string
	LoginPath   string
	SendMsgPath string
	Username    string
	Password    string
}

type CliConfig struct {
	System  *int
	Node    *int
	Host    *string
	Port    *int
	Cfgpath *string
}

type Database struct {
	Mongo *MongoDb
	Redis *RedisDb
}

type MongoDb struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisDb struct {
	Host     string
	Port     string
	Password string
	Database string
}

func LoadConfig() (conf Config, err *error) {
	var cfg CliConfig

	cfg.System = flag.Int("s", 6, "System code as an integer (1-9), 3, 4, 5 etc. default is 6, is testnet")
	cfg.Node = flag.Int("n", 0, "Node code of system, from 0x00-0xFF (001-255), default is 000")
	cfg.Host = flag.String("h", "localhost", "Domain or ip of host, default is localhost")
	cfg.Port = flag.Int("p", 80, "Port of service, default is 80")
	cfg.Cfgpath = flag.String("c", "conf/conf.ini", "Path of program's configuration path, default is conf/conf.ini")
	flag.Parse()
	conf.Environment = &cfg
	readConfigFile(&conf)

	return conf, err
}

func readConfigFile(c *Config) {
	//var err error
	cfg, err := ini.Load(*c.Environment.Cfgpath)
	if err != nil {
		logger.LogErr.Println(err)
		fmt.Println(err)
		os.Exit(0)
	}

	db := &Database{}
	mongodb := cfg.Section("mongodb")
	fmt.Println("Da goi file ini", mongodb.Key("host").String())
	db.Mongo = &MongoDb{Host: mongodb.Key("host").String(), Port: mongodb.Key("port").String(), Username: mongodb.Key("username").String(), Password: mongodb.Key("password").String(), Database: mongodb.Key("database").String()}

	reddb := cfg.Section("redisdb")
	db.Redis = &RedisDb{Host: reddb.Key("host").String(), Port: reddb.Key("port").String(), Password: reddb.Key("password").String(), Database: reddb.Key("database").String()}

	vendor := cfg.Section("smsvendor")
	c.Vendor = &Vendor{Url: vendor.Key("url").String(), LoginPath: vendor.Key("login").String(), SendMsgPath: vendor.Key("login").String(), Username: vendor.Key("username").String(), Password: vendor.Key("password").String()}

	c.Database = db
}
