package dao

import (
	. "apiProject/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "users"

// Find all users
func (u *UserDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Establish a connection to database
func (u *UserDAO) Connect() {
	session, err := mgo.Dial(u.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(u.Database)
}
