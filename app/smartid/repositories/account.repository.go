package repositories

import (
	"fmt"

	"github.com/thiepwong/smartid/app/smartid/models"
	"gopkg.in/mgo.v2"
)

type AccountRepository interface {
	Save(*models.AccountModel) (*models.AccountModel, error)
	Update(*models.AccountModel) (*models.AccountModel, error)
	FindById(uint64) (*models.AccountModel, error)
	FindByUsername(*models.Username) (*models.AccountModel, error)
	FindAll() (*models.AccountModel, error)
}

type accountRepositoryContext struct {
	db         *mgo.Database
	collection string
}

func NewAccountRepositoryContext(db *mgo.Database, collection string) *accountRepositoryContext {
	return &accountRepositoryContext{
		db:         db,
		collection: collection,
	}
}

//Save
func (r *accountRepositoryContext) Save(accountModel *models.AccountModel) (*models.AccountModel, error) {
	fmt.Println("insert vao collect ", r.collection)
	err := r.db.C(r.collection).Insert(accountModel)
	return accountModel, err
}

// Update
func (r *accountRepositoryContext) Update(accountModel *models.AccountModel) (*models.AccountModel, error) {
	return nil, nil
}

//FindbyId
func (r *accountRepositoryContext) FindById(uint64) (*models.AccountModel, error) {
	return nil, nil
}

// FindByUsername
func (r *accountRepositoryContext) FindByUsername(accountModel *models.Username) (*models.AccountModel, error) {
	var acc models.AccountModel
	r.db.C(r.collection).Find(accountModel).One(&acc)
	return &acc, nil
}

//FindAll
func (r *accountRepositoryContext) FindAll() (*models.AccountModel, error) {
	return nil, nil
}
