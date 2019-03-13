package services

import (
	"encoding/json"
	"fmt"

	"github.com/thiepwong/smartid/app/smartid/models"
	"github.com/thiepwong/smartid/app/smartid/repositories"
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
	account := models.SignupModel{ID: 3435353535445,
		Username: models.Username{Mobile: "0983851116",
			Email: "thiep.wong@gmail.com"},
		Mobile:   "0983851116",
		Email:    "thiep.wong@gmail.com",
		Fulname:  "Hoang Van Thiep",
		Birthday: 12312312313,
		Profile:  models.Profiles{Avatar: "http://avatar.com", Cover: "http://cover.com"}}
	//myModel.CreatedTime = time.Now()
	//return db.Insert(s.Table, myModel)
	err := s.Table.Insert(account)
	if err != nil {
		fmt.Println("Loi insert roi")
		return err.Error()
	}

	j, er := json.Marshal(account)

	if er != nil {
		return "error!"
	}

	return string(j)
}

func (s *accountService) Signin(username string, password string) bool {
	return true
}
