package repositories

import "github.com/thiepwong/smartid/app/smartid/models"

type Query func(models.SignupModel) bool
type AccountRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (signup models.SignupModel, found bool)
	SelectMany(query Query, limit int) (results []models.SignupModel)
	Get() string

	InsertOrUpdate(signup models.SignupModel) (updatedMovie models.SignupModel, err error)
	Delete(query Query, limit int) (deleted bool)
}

func Get() string {
	return "Da goi vao trong repo roi"
}
