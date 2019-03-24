package repositories

import (
	"time"

	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/thiepwong/smartid/pkg/datasource/redis"
)

type OtpRepository interface {
	Save(string, interface{}, time.Duration) bool
	Read(string) string
	Delete(string) bool
}

type otpRepositoryImp struct {
	redis *redis.Redis
}

func NewOtpRepository(redis *redis.Redis) OtpRepository {
	return &otpRepositoryImp{redis}
}

func (r *otpRepositoryImp) Delete(key string) bool {
	return true
}

func (r *otpRepositoryImp) Read(key string) string {
	val, err := r.redis.Client.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (r *otpRepositoryImp) Save(key string, value interface{}, ttl time.Duration) bool {
	err := r.redis.Client.Set(key, value, ttl*1000000000).Err()
	if err != nil {
		logger.LogErr.Println(err.Error())
		return false
	}

	return true
}
