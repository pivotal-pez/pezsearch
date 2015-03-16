package pezinventory

import "github.com/go-martini/martini"

//InitRoutes - initialize the mappings for controllers against valid routes
func InitRoutes(m *martini.ClassicMartini) {

	//types routes
	m.Group("/v1/types", func(r martini.Router) {
		ctrl := &TypeController{}
		r.Get("", ctrl.ListTypes)
		r.Get("/:id", ctrl.GetType)
		r.Get("/:id/items", ctrl.ListTypeItems)
	})

	//items routes
	m.Group("/v1/items", func(r martini.Router) {
		ctrl := &ItemController{}
		r.Get("", ctrl.ListItems)
		r.Get("/:id", ctrl.GetItem)
		r.Get("/:id/history", ctrl.GetItemHistory)
	})
}
