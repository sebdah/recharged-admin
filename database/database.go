package database

import (
	"log"

	"github.com/sebdah/recharged-admin/config"
	"github.com/sebdah/recharged-admin/database"
	"github.com/sebdah/recharged-admin/models"
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

// Ensure databases
func EnsureDatabases() {
	if config.Env == "dev" {
		log.Println("Ensuring databases")
		database.CreateCollectionIdTags()
		database.CreateCollectionChargePoints()
	}
}

// Ensure indexes
func EnsureIndexes() {
	if config.Env == "dev" {
		log.Println("Ensuring indexes")
		models.EnsureIndexes(new(models.IdTag))
		models.EnsureIndexes(new(models.ChargePoint))
	}
}
