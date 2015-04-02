package pezsearch

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

//Server wraps the Martini server struct
type Server *martini.ClassicMartini

//NewServer configures and returns a Server.
//TODO: Parameterize DB
func NewServer() Server {

	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))

	//TODO: Database

	m.Group("/v1/search", func(r martini.Router) {
		searchCtrl := searchController{}
		r.Get("", searchCtrl.find)

	})

	return m
}
