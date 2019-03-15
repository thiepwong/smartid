package repositories

import (
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

type AccountRepositoryContext struct {
	db         *mgo.Database
	collection string
}

func NewAccountRepositoryContext(db *mgo.Database, collection string) *AccountRepositoryContext {
	return &AccountRepositoryContext{
		db:         db,
		collection: collection,
	}
}

//Save
func (r *AccountRepositoryContext) Save(accountModel *models.AccountModel) (*models.AccountModel, error) {
	err := r.db.C(r.collection).Insert(accountModel)
	return accountModel, err
}

// Update
func (r *AccountRepositoryContext) Update(accountModel *models.AccountModel) (*models.AccountModel, error) {
	return nil, nil
}

//FindbyId
func (r *AccountRepositoryContext) FindById(uint64) (*models.AccountModel, error) {
	return nil, nil
}

// FindByUsername
func (r *AccountRepositoryContext) FindByUsername(accountModel *models.Username) (*models.AccountModel, error) {
	return nil, nil
}

//FindAll
func (r *AccountRepositoryContext) FindAll() (*models.AccountModel, error) {
	return nil, nil
}
