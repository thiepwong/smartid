package repositories

type SmsRepository interface {
	Save(string)
	Read() string
	Delete(string) bool
}

type smsRepositoryImp struct {
}
