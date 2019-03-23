package repositories

import (
	"time"

	"github.com/thiepwong/smartid/pkg/datasource/redis"
)

type SmsRepository interface {
	Save(string, interface{}, time.Duration) bool
	Read(string) string
	Delete(string) bool
}

type smsRepositoryImp struct {
	redis *redis.Redis
}

func NewSmsRepository(redis *redis.Redis) SmsRepository {
	return &smsRepositoryImp{redis}
}

func (r *smsRepositoryImp) Delete(key string) bool {
	return true
}

func (r *smsRepositoryImp) Read(key string) string {
	val, err := r.redis.Client.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (r *smsRepositoryImp) Save(key string, value interface{}, ttl time.Duration) bool {
	err := r.redis.Client.Set(key, value, ttl*1000000000).Err()
	if err != nil {
		return false
	}

	return true
}
