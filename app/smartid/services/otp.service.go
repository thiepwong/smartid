package services

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/thiepwong/smartid/app/smartid/repositories"
	"github.com/thiepwong/smartid/pkg/logger"
)

type OtpService interface {
	SendOTP(string, time.Duration) bool
}

type otpServiceImp struct {
	Repo repositories.OtpRepository
}

type OTP struct {
	Code   string        `json:"code"`
	Mobile string        `json:"mobile"`
	TTL    time.Duration `json:"ttl"`
}

//NewOtpService func
func NewOtpService(repo repositories.OtpRepository) OtpService {
	return &otpServiceImp{repo}
}

func (s *otpServiceImp) SendOTP(mobile string, ttl time.Duration) bool {
	_otp := generateOTP(mobile, 6, ttl)
	_json, err := json.Marshal(_otp)
	if err != nil {

		logger.LogErr.Println(err.Error())
		return false
	}
	// using sms microservice to send this otp

	

	return s.Repo.Save(_otp.Code, string(_json), ttl)
}

func generateOTP(mobile string, size int, ttl time.Duration) *OTP {
	var code string
	for i := 0; i < size; i++ {
		code += strconv.Itoa(rand.Intn(9))
	}
	_otp := &OTP{Code: code, Mobile: mobile, TTL: ttl}
	return _otp
}

func deleteOTP(string) bool {
	return true
}
