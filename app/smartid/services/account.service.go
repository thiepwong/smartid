package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/thiepwong/smartid/pkg/wallet"

	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/repositories"
	"github.com/thiepwong/smartid/pkg/luhn"
	"gopkg.in/mgo.v2"
)

type AccountService interface {
	Signup(data *models.SignupModel) bool
	Signin(username string, password string) bool
	Get() string
}

type accountService struct {
	//	repo  repositories.AccountRepository
	Table *mgo.Collection
}

//SignupService signup a service
func NewAccountService(db *mgo.Database, session *mgo.Session) AccountService {
	return &accountService{Table: db.C("accounts")}

}

//RegSignupService   Register a new Service
func RegSignupService(repo repositories.AccountRepository) AccountService {
	return &accountService{}
}

func (s *accountService) Signup(data *models.SignupModel) bool {
	return true
}

func (s *accountService) Get() string {

	_begin := time.Now()
	fmt.Println("Bat dau: ", _begin)
	// var _accArray []models.AccountModel
	// for i := 0; i < 10000; i++ {

	_id, err := luhn.GenerateSmartID(8, 0x8, 16)
	_wl, err := wallet.CreateWallet()

	//	fmt.Println("Da tao duoc vi la: ", _wl.Address, _wl.PublicKey)
	_acModel := models.AccountModel{ID: _id,
		Username:  models.Username{Mobile: strconv.FormatUint(_id, 10), Email: "abc@" + strconv.FormatUint(_id, 10)},
		Mobile:    strconv.FormatUint(_id, 10),
		Email:     "abc@" + strconv.FormatUint(_id, 10),
		Firstname: "Hoang",
		Midname:   "Van",
		Lastname:  "Thiep",
		SocialID:  []models.SocialID{models.SocialID{Network: "facebook", Id: "345345353453"}},
		Wallet:    _wl,
		Profiles:  models.Profiles{Avatar: "http://avatar.com", Cover: "http://cover.com"}}

	//_accArray = append(_accArray, _acModel)
	//	}

	err = s.Table.Insert(_acModel)
	if err != nil {
		fmt.Println("Loi insert roi")
		return err.Error()
	}

	_begin.Unix()
	_end := time.Now()
	fmt.Println("Ket thuc: ", _end)
	fmt.Println("Thoi gian thuc hien: ", _end.Unix()-_begin.Unix())
	j, er := json.Marshal(testTime{begin: _begin, end: _end})

	if er != nil {
		return "error!"
	}

	return string(j)
}

type testTime struct {
	begin time.Time
	end   time.Time
}

func (s *accountService) Signin(username string, password string) bool {
	return true
}
