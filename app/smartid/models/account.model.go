package models

// SignupModel struct is member signup data
type SignupModel struct {
	ID       int64 `bson:"_id,omitempty"`
	Username Username
	Mobile   string
	Email    string
	Fulname  string
	Birthday int
	Profile  Profiles
}

type Username struct {
	Mobile string
	Email  string
}

type Profiles struct {
	Cover    string
	Avatar   string
	Nickname string
}
