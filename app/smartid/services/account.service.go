package services

import (
	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/repositories"
)

type AccountService interface {
	Signup(data *models.SignupModel) bool
	Signin(username string, password string) bool
	Get() string
}

type accountService struct {
	repo repositories.AccountRepository
}

func SignupService(repo repositories.AccountRepository) AccountService {
	return &accountService{repo: repo}

}

func RegSignupService() AccountService {
	return &accountService{}
}

func (s *accountService) Signup(data *models.SignupModel) bool {
	return true
}

func (s *accountService) Get() string {
	//	account := models.SignupModel{Mobile: "0983851116", Email: "1234"}
	return s.repo.Get()
}

func (s *accountService) Signin(username string, password string) bool {
	return true
}
