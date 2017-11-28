package storage

import (
	"log"

	"gopkg.in/mgo.v2"
)

func Dbinit() *mgo.Session {
	session, err := mgo.Dial("localhost")
	log.Printf("created session")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
