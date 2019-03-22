package datasource

import (
	"fmt"
	"strconv"

	"github.com/thiepwong/smartid/pkg/config"
	"github.com/thiepwong/smartid/pkg/datasource/redis"
)

func GetRedisDb(cfg *config.RedisDb) *redis.Redis {
	var _red redis.Redis
	var e error
	_red.Options.Addr = cfg.Host + ":" + cfg.Port
	_red.Options.Password = cfg.Password
	_red.Options.DB, e = strconv.Atoi(cfg.Database)
	if e != nil {
		_red.Options.DB = 0
	}
	_re := redis.RegisterRedis(_red)
	fmt.Println(cfg.Host, cfg.Port, cfg.Password, cfg.Database)
	return &_re

}
