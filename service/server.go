package pezsearch

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

//Server wraps the Martini server struct
type Server *martini.ClassicMartini

//NewServer configures and returns a Server.
//TODO: Parameterize DB
func NewServer(authHandler martini.Handler) (m Server) {

	m = Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))

	m.Use(cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
		AllowHeaders:    []string{"X-API-KEY"},
	}))

	//TODO: Database
	m.Group("/v1/search", func(r martini.Router) {
		searchCtrl := searchController{}
		r.Get("", searchCtrl.find)

	}, authHandler)
	return
}
