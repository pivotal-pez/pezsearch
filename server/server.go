package server

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/pivotalservices/pezinventory/controllers"
)

//Server wraps the Martini server struct
type Server *martini.ClassicMartini

func NewServer() Server {

	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))

	//TODO: Database

	//types routes
	m.Group("/v1/types", func(r martini.Router) {
		ctrl := &controllers.TypeController{}
		r.Get("", ctrl.ListTypes)
		r.Get("/:id", ctrl.GetType)
		r.Get("/:id/items", ctrl.ListTypeItems)
	})

	//items routes
	m.Group("/v1/items", func(r martini.Router) {
		ctrl := &controllers.ItemController{}
		r.Get("", ctrl.ListItems)
		r.Get("/:id", ctrl.GetItem)
		r.Get("/:id/history", ctrl.GetItemHistory)
	})

	return m
}
