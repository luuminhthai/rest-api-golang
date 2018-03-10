package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	firstName string        `form:"firstName" json:"firstName"`
	lastName  string        `form:"lastName" json:"lastName"`
	email     string        `form:"email" json:"email"`
	address   string        `form:"email" json:"address"`
}
