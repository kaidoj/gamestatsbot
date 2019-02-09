package models

import (
	"testing"

	"github.com/globalsign/mgo"
	"github.com/niilo/golib/test/dockertest"
)

func TestFindByID(t *testing.T) {
	containerID, ip := dockertest.SetupMongoContainer(t)
	defer containerID.KillRemove(t)

	mongoSession, err := mgo.Dial(ip)
	if err != nil {
		t.Errorf("MongoDB connection failed, with address")
	}

	db := &DB{mongoSession, nil}
	ok, err := db.FindByID("123123")
	if err != nil || ok == nil {
		t.Errorf("No user found with id 123123")
	}
}
