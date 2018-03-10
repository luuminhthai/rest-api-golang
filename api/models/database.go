package models

import (
	"os"

	"gopkg.in/mgo.v2"
)

func Database() *mgo.Session {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	session.DB(os.Getenv("DB_NAME"))
	return session
}
