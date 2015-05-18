package database

import (
	"github.com/sebdah/recharged-admin/config"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

// Connect to MongoDB
func GetSession() *mgo.Session {
	if Session == nil {
		session, err := mgo.Dial(config.Config.GetString("mongodb.hosts"))
		if err != nil {
			panic(err)
		}
		Session = session
	}

	return Session
}

// Get the database
func GetDb() *mgo.Database {
	if Session == nil {
		GetSession()
	}

	return Session.DB(config.Config.GetString("mongodb.db"))
}
