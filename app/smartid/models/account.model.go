package models

import "github.com/thiepwong/smartid/pkg/wallet"

// SignupModel struct is member signup data
type SignupModel struct {
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Birthday  int
	Gender    int
}

type SigninModel struct {
	Username string
	Password string
}

type AccountModel struct {
	ID        uint64 `bson:"_id,omitempty"`
	Wallet    wallet.Wallet
	Username  Username
	Mobile    string
	Email     string
	Firstname string
	Lastname  string
	Birthday  int
	SocialID  []SocialID
	Profiles  Profiles
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

type SocialID struct {
	Network string
	Id      string
}
