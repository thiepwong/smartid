package datasources

import (
	"gopkg.in/mgo.v2"
)

func GetMongoDb() (*mgo.Database, error) {
	// host := os.Getenv("HOST_MONGO")
	dbName := "test"

	// session, err := mgo.Dial(host)

	dialInfo, err := mgo.ParseURL("mongodb://171.244.49.164:2688/ucenter")

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		return nil, err
	}
	db := session.DB(dbName)

	return db, nil
}

func GetSession() *mgo.Session {
	dialInfo, err := mgo.ParseURL("mongodb://171.244.49.164:2688/ucenter")
	s, e := mgo.DialWithInfo(dialInfo)
	if e != nil {
		panic(err)
	}
	return s
}
