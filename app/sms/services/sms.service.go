package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url "net/url"
	"strconv"
	"time"

	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/thiepwong/smartid/app/sms/repositories"
	"github.com/thiepwong/smartid/pkg/config"
)

// SmsService interface of service
type SmsService interface {
	Login() interface{}
	SendMsg(string, string) interface{}
}

type smsServiceImp struct {
	smsRepo repositories.SmsRepository
	vendor  *config.Vendor
}

type smsToken struct {
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	ExpiresIn   time.Duration `json:"expires_in"`
}

//NewSmsService function for register new service
func NewSmsService(repo repositories.SmsRepository, cfg *config.Vendor) SmsService {
	return &smsServiceImp{repo, cfg}
}

func (s *smsServiceImp) Login() interface{} {
	_token := vendorLogin(s.vendor.Url+s.vendor.LoginPath, s.vendor.Username, s.vendor.Password)
	_smsToken := smsToken{}
	json.Unmarshal([]byte(_token), &_smsToken)
	s.smsRepo.Save("sms_token", _token, _smsToken.ExpiresIn)
	return &_smsToken
}

func (s *smsServiceImp) SendMsg(mobile string, msg string) interface{} {
	_smsToken := smsToken{}
	_token := s.smsRepo.Read("sms_token")
	if _token == "" {
		_token = vendorLogin(s.vendor.Url+s.vendor.LoginPath, s.vendor.Username, s.vendor.Password)
		json.Unmarshal([]byte(_token), &_smsToken)
		s.smsRepo.Save("sms_token", _token, _smsToken.ExpiresIn)

	}

	json.Unmarshal([]byte(_token), &_smsToken)
	return sendMessage(s.vendor.Url+s.vendor.SendMsgPath, _smsToken.AccessToken, mobile, msg)

}

func vendorLogin(path string, username string, password string) string {
	auth := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	client := &http.Client{}

	data := url.Values{}
	data.Add("grant_type", "password")

	req, err := http.NewRequest("POST", path, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+encoded)
	//	req.PostForm.Add("grant_type", "password")
	resp, err := client.Do(req)
	if err != nil {
		logger.LogErr.Println(err)
		return ""
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(f)

}

func sendMessage(path string, token string, mobile string, msg string) interface{} {
	client := &http.Client{}
	data := url.Values{}
	_bodyStr := fmt.Sprintf(`{
		"Brandname": "SMARTLIFE",
		"SendingList": [
			{
				"SmsId": "SMS_%s"  ,
				"PhoneNumber": "%s",
				"Content":  "%s",
				"ContentType": "30"
			}
		]
	}`, strconv.Itoa(int(time.Now().Unix())), mobile, msg)

	js := []byte(_bodyStr)
	data.Add("Brandname", "SMARTLIFE")
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(js))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		logger.LogErr.Println(err)
		return ""
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(f)
}
