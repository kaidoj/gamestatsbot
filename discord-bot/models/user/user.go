package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/kaidoj/gamestatsbot/discord-bot/db"
	models "github.com/kaidoj/gamestatsbot/discord-bot/models/model"
)

var userCollection string = "users"

//User structure
type User struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	models.ModelImpl
	UserID       string
	Username     string
	MessageCount int
}

//FindByID gets user by id
func (u *User) FindByID(userID string) (*User, error) {
	session := db.Con.GetSession().Copy()
	defer session.Close()
	c := db.Con.GetDb(session).C(userCollection)

	user := &User{}
	err := c.Find(bson.M{"userid": userID}).One(&user)
	return user, err
}

//Insert user to db
func (u *User) Insert() error {
	session := db.Con.GetSession().Copy()
	defer session.Close()
	c := db.Con.GetDb(session).C(userCollection)

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	err := c.Insert(u)
	return err
}

//Update user to db
func (u *User) Update() error {
	session := db.Con.GetSession().Copy()
	defer session.Close()
	c := db.Con.GetDb(session).C(userCollection)
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
func ResetMessagesCount() error {
	session := db.Con.GetSession().Copy()
	defer session.Close()
	c := db.Con.GetDb(session).C(userCollection)
	colQuerier := bson.M{}
	change := bson.M{"$set": bson.M{
		"messagecount": 0,
		"updatedat":    time.Now(),
	}}
	err := c.Update(colQuerier, change)
	return err
}

//CreateOrUpdate user
func (u *User) CreateOrUpdate() error {
	var err error
	userExists, err := u.FindByID(u.UserID)
	if err == nil {
		userExists.MessageCount = u.MessageCount
		userExists.Username = u.Username
		err = userExists.Update()
	} else {
		err = u.Insert()
	}

	return err
}

//GetUsersByMessageCount users by message count
func (u *User) GetUsersByMessageCount(limit int) []User {
	session := db.Con.GetSession().Copy()
	defer session.Close()
	c := db.Con.GetDb(session).C(userCollection)

	var results []User
	err := c.Find(bson.M{}).Sort("-messagecount").Limit(limit).All(&results)
	if err != nil {
		log.Printf("GetUsersByMessageCount() could not fetch messages: %v", err)
	}

	return results
}

//UpdateMessageCount updates message count for user
func (u *User) UpdateMessageCount() {
	user, err := u.FindByID(u.UserID)
	if err == nil {
		user.MessageCount++
		user.Update()
	} else {
		u.MessageCount = 1
		u.Insert()
	}
}
