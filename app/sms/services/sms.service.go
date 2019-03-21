package services

import (
	"github.com/thiepwong/smartid/app/sms/repositories"
)

type SmsService interface {
	Login(username string, password string) string
	SendSms(message string) string
}

type smsServiceImp struct {
	smsRepo repositories.SmsRepository
}

func NewSmsService(repo repositories.SmsRepository) SmsService {
	return &smsServiceImp{repo}
}

func (s *smsServiceImp) Login(username string, password string) string {
	return ""
}

func (s *smsServiceImp) SendSms(msg string) string {
	return ""
}
