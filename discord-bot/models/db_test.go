package models

import (
	"os"
	"path"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/kaidoj/gamestatsbot/discord-bot/config"
)

func TestGetSession(t *testing.T) {
	db := &DB{&mgo.Session{}, nil}
	session := db.GetSession()
	if session == nil {
		t.Errorf("Database session not found")
	}
}

func TestGetDb(t *testing.T) {
	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/kaidoj/gamestatsbot/discord-bot")
	db := &DB{&mgo.Session{}, config.Init(ap)}
	database := db.GetDb(db.GetSession())
	if database == nil {
		t.Errorf("No database returned")
	}
}

func TestGetCopy(t *testing.T) {
	db := &DB{nil, nil}
	copy := db.GetCopy()
	if copy == nil {
		t.Errorf("No database copy returned")
	}
}
