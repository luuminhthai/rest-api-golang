package main

import (
	"github.com/go-martini/martini"
	"github.com/luuminhthai/rest-api-golang/api/controllers"
	"github.com/luuminhthai/rest-api-golang/api/models"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Map(models.Database())
	m.Use(render.Renderer())

	pc := controllers.newUserController(models.Database())

	m.Get("/users", binding.Bind(models.User{}), pc.GetAllUsers)
	m.Post("/users", binding.Bind(models.User{}), pc.PostUser)
	m.Run()
}
