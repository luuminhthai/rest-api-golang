package controllers

import (
	"os"

	"github.com/luuminhthai/rest-api-golang/api/models"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (pc *UserController) GetAllUsers(r render.Render) {
	users := []models.User{}
	session := pc.session.DB(os.Getenv("DB_NAME")).C("users")
	err := session.Find(nil).Limit(100).All(&users)

	if err != nil {
		panic(err)
	}

	r.JSON(200, users)
}

// func (pc *UserController) PostUser(user models.User, r render.Render) {
// 	session := pc.session.DB(os.Getenv("DB_NAME")).C("users")

// 	user.Id = bson.NewObjectId()
// 	user.firstName = user.firstName
// 	user.lastName = user.lastName
// 	user.address = user.address
// 	user.email = user.email
// 	session.Insert(user)
// 	r.JSON(200, user)
// }
