package services

import (
	"github.com/thiepwong/smartid/app/sms/repositories"
	"github.com/thiepwong/smartid/pkg/config"
)

type SmsService interface {
	Login(username string, password string) string
	SendSms(message string) string
}

type smsServiceImp struct {
	smsRepo repositories.SmsRepository
	vendor  *config.Vendor
}

func NewSmsService(repo repositories.SmsRepository, cfg *config.Vendor) SmsService {
	return &smsServiceImp{repo, cfg}
}

func (s *smsServiceImp) Login(username string, password string) string {

	//s.vendor.
	s.smsRepo.Save(username, password, 200000)
	c := s.smsRepo.Read(username)
	return "Da ghi xong " + c
}

func (s *smsServiceImp) SendSms(msg string) string {
	return s.smsRepo.Read(msg)
}
