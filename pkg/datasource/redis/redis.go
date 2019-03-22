package redis

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	Options redis.Options
	Client  *redis.Client
}

func RegisterRedis(red Redis) Redis {
	red.Client = redis.NewClient(&red.Options)
	return red
}
