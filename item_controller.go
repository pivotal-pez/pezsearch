package pezsearch

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type itemController struct {
}

func (c *itemController) listItems(render render.Render) {
	i := listItems()
	render.JSON(200, successMessage(i))
}

func (c *itemController) getItem(params martini.Params, render render.Render) {
	i := getItem(params["id"])
	render.JSON(200, successMessage(i))
}

func (c *itemController) getItemHistory(params martini.Params, render render.Render) {
	render.JSON(501, errorMessage("Not Implemented"))
}
