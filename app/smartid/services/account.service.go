package services

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/repositories"
	"github.com/thiepwong/smartid/pkg/luhn"
	"github.com/thiepwong/smartid/pkg/wallet"
)

type AccountService interface {
	RegisterAccount(userData *models.SignupModel) (*models.AccountModel, error)
	UpdateAccount(profile *models.AccountModel) (*models.AccountModel, error)
	ActiveAccount(id uint64) (bool, error)
	SigninAccount(signinData models.SigninModel) (string, error)
	Test(profile *models.SignupModel) string
}

type accountServiceImpl struct {
	accountRepository repositories.AccountRepository
}

//NewAccountService register new Service
func NewAccountService(repo repositories.AccountRepository) AccountService {

	return &accountServiceImpl{accountRepository: repo}
}

// RegisterAccount
func (accountService *accountServiceImpl) RegisterAccount(account *models.SignupModel) (*models.AccountModel, error) {

	// fmt.Println("Ten da nhap", account.Username)
	var _acModel models.AccountModel
	var Err error
	_acModel.Username, _acModel.Mobile, _acModel.Email, Err = validateUsername(account.Username)
	if Err != nil {
		return nil, Err
	}

	_isExist := checkExist(accountService, &_acModel.Username)

	if _isExist == true {
		return nil, errors.New("username is existed")
	}

	_id, err := luhn.GenerateSmartID(8, 0x8, 16)
	_wl, err := wallet.CreateWallet()

	if err != nil {
		log.Fatal(err)
	}

	_acModel.ID = _id

	_acModel.Birthday = account.Birthday
	_acModel.Firstname = account.Firstname
	_acModel.Lastname = account.Lastname
	_acModel.Wallet = _wl
	res, err := accountService.accountRepository.Save(&_acModel)
	return res, err
}

//Signin
func (accountService *accountServiceImpl) Signin(userInfo models.SigninModel) string {
	//	b, _ := json.Marshal(*userInfo)
	//	return string(b)
	fmt.Print("Da vao trong nay roi")
	return "da vao trong service!"
}

//Test func
func (s *accountServiceImpl) Test(profile *models.SignupModel) string {
	return "Da doc duoc trong server la: " + profile.Firstname
}

func (s *accountServiceImpl) ActiveAccount(id uint64) (bool, error) {
	return true, nil
}

func (s *accountServiceImpl) SigninAccount(signinData models.SigninModel) (string, error) {
	return "da dang nhap", nil
}

func (s *accountServiceImpl) UpdateAccount(profile *models.AccountModel) (*models.AccountModel, error) {
	return nil, nil
}

func validateUsername(username string) (Username models.Username, Mobile string, Email string, Err error) {
	if username == "" {
		Err = errors.New("Invalid username")
		return Username, Mobile, Mobile, Err
	}
	emailReg := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isEmail := emailReg.MatchString(username)
	if isEmail == true {
		Email = username
		Username.Email = username
	} else {

		mobileReg := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
		isMobile := mobileReg.MatchString(username)
		if isMobile == true {
			Mobile = username
			Username.Mobile = username
		} else {
			Err = errors.New("Invalide Username")
		}

	}

	return Username, Mobile, Email, Err

}

func checkExist(s *accountServiceImpl, username *models.Username) bool {

	a, e := s.accountRepository.FindByUsername(username)
	if e != nil {
		logger.LogErr.Println(e)
		return false
	}

	if a.ID != 0 {
		return true
	} else {
		return false
	}

}
