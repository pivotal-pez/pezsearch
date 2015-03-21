package pezsearch

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type typeController struct {
}

func (c *typeController) listTypes(render render.Render) {
	t := listTypes()
	render.JSON(200, successMessage(t))
}

func (c *typeController) getType(params martini.Params, render render.Render) {
	t := getType(params["id"])
	render.JSON(200, successMessage(t))
}

func (c *typeController) listTypeItems(params martini.Params, render render.Render) {
	i := listItemsByType(params["id"])
	render.JSON(200, successMessage(i))
}
