package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

const userCollection = "users"

type UserInterface interface {
	GetUsersByMessageCount(limit int) []*User
	CreateOrUpdate(u *User) error
	ResetMessagesCount() error
	UpdateMessageCount(u *User) error
}

//User structure
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	UserID       string
	Username     string
	MessageCount int
	Model
}

//FindByID gets user by id
func (db *DB) FindByID(userID string) (*User, error) {
	session := db.GetSession().Copy()
	defer session.Close()
	c := db.GetDb(session).C(userCollection)

	user := &User{}
	err := c.Find(bson.M{"userid": userID}).One(&user)
	return user, err
}

//Insert user to db
func (db *DB) Insert(u *User) error {
	session := db.GetSession().Copy()
	defer session.Close()
	c := db.GetDb(session).C(userCollection)

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	err := c.Insert(u)
	return err
}

//Update user to db
func (db *DB) Update(u *User) error {
	session := db.GetSession().Copy()
	defer session.Close()
	c := db.GetDb(session).C(userCollection)
	colQuerier := bson.M{"_id": u.ID}
	change := bson.M{"$set": bson.M{
		"userid":       u.UserID,
		"messagecount": u.MessageCount,
		"updatedat":    time.Now(),
		"username":     u.Username,
	}}
	err := c.Update(colQuerier, change)
	return err
}

//ResetMessagesCount reset user messages count
func (db *DB) ResetMessagesCount() error {
	session := db.GetSession().Copy()
	defer session.Close()
	c := db.GetDb(session).C(userCollection)
	colQuerier := bson.M{}
	change := bson.M{"$set": bson.M{
		"messagecount": 0,
		"updatedat":    time.Now(),
	}}
	err := c.Update(colQuerier, change)
	return err
}

//CreateOrUpdate user
func (db *DB) CreateOrUpdate(u *User) error {
	var err error
	userExists, err := db.FindByID(u.UserID)
	if err == nil {
		userExists.MessageCount = u.MessageCount
		userExists.Username = u.Username
		err = db.Update(userExists)
	} else {
		err = db.Insert(u)
	}

	return err
}

//GetUsersByMessageCount users by message count
func (db *DB) GetUsersByMessageCount(limit int) []*User {
	session := db.GetSession().Copy()
	defer session.Close()
	c := db.GetDb(session).C(userCollection)

	var results []*User
	err := c.Find(bson.M{}).Sort("-messagecount").Limit(limit).All(&results)
	if err != nil {
		log.Printf("GetUsersByMessageCount() could not fetch messages: %v", err)
	}

	return results
}

//UpdateMessageCount updates message count for user
func (db *DB) UpdateMessageCount(u *User) error {
	user, err := db.FindByID(u.UserID)
	var failed error
	if err == nil {
		user.MessageCount++
		failed = db.Update(u)
	} else {
		u.MessageCount = 1
		failed = db.Insert(u)
	}

	return failed
}
