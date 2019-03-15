package datasources

import (
	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {
	dialInfo, err := mgo.ParseURL("mongodb://171.244.49.164:2688/ucenter")
	s, e := mgo.DialWithInfo(dialInfo)
	if e != nil {
		panic(err)
	}
	return s
}
