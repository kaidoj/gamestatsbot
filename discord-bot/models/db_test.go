package models

import (
	"os"
	"path"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/kaidoj/gamestatsbot/discord-bot/config"
	"github.com/spf13/viper"
)

func initConfig() *viper.Viper {
	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/kaidoj/gamestatsbot/discord-bot")

	return config.Init(ap)
}

func TestInit(t *testing.T) {
	database := &DB{&mgo.Session{}, initConfig()}
	db := database.Init()

	if db == nil {
		t.Errorf("Database connection not enstablished")
	}
}

func TestGetSession(t *testing.T) {
	database := &DB{&mgo.Session{}, initConfig()}
	db := database.Init()
	session := db.GetSession()

	if session == nil {
		t.Errorf("Session not found")
	}
}
