package config

import "flag"

type Config struct {
	System  *int
	Node    *int
	Host    *string
	Port    *int
	Cfgpath *string
}

func LoadConfig() (cfg Config, err *error) {

	cfg.System = flag.Int("s", 6, "System code as an integer (1-9), 3, 4, 5 etc. default is 6, is testnet")
	cfg.Node = flag.Int("n", 0, "Node code of system, from 0x00-0xFF (001-255), default is 000")
	cfg.Host = flag.String("h", "localhost", "Domain or ip of host, default is localhost")
	cfg.Port = flag.Int("p", 80, "Port of service, default is 80")
	cfg.Cfgpath = flag.String("c", "conf/conf.ini", "Path of program's configuration path, default is conf/conf.ini")
	flag.Parse()
	return cfg, err
}
