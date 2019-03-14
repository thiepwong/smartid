package models

import "github.com/thiepwong/smartid/pkg/wallet"

// SignupModel struct is member signup data
type SignupModel struct {
	ID       uint64 `bson:"_id,omitempty"`
	Username Username
	Mobile   string
	Email    string
	Fulname  string
	Birthday int
	Profile  Profiles
}

type AccountModel struct {
	ID        uint64 `bson:"_id,omitempty"`
	Wallet    wallet.Wallet
	Username  Username
	Mobile    string
	Email     string
	Firstname string
	Midname   string
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
