package models

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

type ConnectionInterface interface {
	GetSession() *mgo.Session
	GetDb(session *mgo.Session) *mgo.Database
	GetCopy() *DB
	UserInterface
}

type DB struct {
	Session *mgo.Session
	Config  *viper.Viper
}

func (db *DB) Init() *DB {
	session, err := mgo.Dial(db.Config.GetString("mongo_url"))
	if err != nil {
		log.Fatalf("New mongo session could not enstablished: %v", err)
	}

	session.Login(&mgo.Credential{
		Username: db.Config.GetString("mongo_user"),
		Password: db.Config.GetString("mongo_pass"),
	})

	db.Session = session

	if db == nil {
		log.Fatalf("Connection to mongodb not enstablished")
	}

	return db
}

func (db *DB) GetSession() *mgo.Session {
	return db.Session
}

func (db *DB) GetDb(session *mgo.Session) *mgo.Database {
	database := session.DB(db.Config.GetString("mongo_db"))
	return database
}

func (db *DB) GetCopy() *DB {
	return db
}
