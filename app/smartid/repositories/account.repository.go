package repositories

import (
	"sync"
)

//type Query func(models.SignupModel) bool
type AccountRepository interface {
	//	Exec(query Query, action Query, limit int, mode int) (ok bool)

	//	Select(query Query) (signup models.SignupModel, found bool)
	//	SelectMany(query Query, limit int) (results []models.SignupModel)
	Get() string

	//	InsertOrUpdate(signup models.SignupModel) (updatedMovie models.SignupModel, err error)
	//	Delete(query Query, limit int) (deleted bool)
}

func (r *accountRepository) Get() string {

	return r.source
}

func RegisterNewAccountRepository(source string) AccountRepository {
	return &accountRepository{source: source}
}

type accountRepository struct {
	source string
	mu     sync.RWMutex
}
