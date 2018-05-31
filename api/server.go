package main

import (
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
)

type User struct {
	firstName string
	lastName  string
}

func main() {
	app := iris.Default()

	session, err := mgo.Dial("localhost:27018")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("testinghi").C("users")
	c.Insert(&User{"luu", "minh thai"})

	app.Get("/users", func(ctx iris.Context) {
		result := User{}
		ctx.JSON(result)
	})

	app.Run(iris.Addr(":8080"))
}
