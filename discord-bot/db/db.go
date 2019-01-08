package db

import (
	"github.com/globalsign/mgo"
	"github.com/kaidoj/gamestatsbot/discord-bot/config"
	"github.com/spf13/viper"
)

var Con *Connection

type Connection struct {
	session *mgo.Session
	db      *mgo.Database
	config  *viper.Viper
}

func init() {
	Con = &Connection{}

	Con.config = config.InitConfig()
	session, err := mgo.Dial(Con.config.GetString("mongo_url"))
	if err != nil {
		panic(err)
	}

	Con.session = session
	session.Login(&mgo.Credential{
		Username: Con.config.GetString("mongo_user"),
		Password: Con.config.GetString("mongo_pass"),
	})

	if Con == nil {
		panic("Connection to mongodb not enstablished")
	}
}

func (con *Connection) GetSession() *mgo.Session {
	return con.session
}

func (con *Connection) GetDb(session *mgo.Session) *mgo.Database {
	db := session.DB(Con.config.GetString("mongo_db"))
	Con.db = db
	return db
}

func (con *Connection) GetCopy() *Connection {
	return con
}
