package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

//ItemController - controller for searching item information
type ItemController struct {
}

func (c *ItemController) ListItems(render render.Render) {
	render.JSON(200, map[string]interface{}{"collection": "List of Items"})
}

func (c *ItemController) GetItem(params martini.Params, render render.Render) {
	render.JSON(200, map[string]interface{}{"item": "Details of Item-" + params["id"]})
}

func (c *ItemController) GetItemHistory(params martini.Params, render render.Render) {
	render.JSON(200, map[string]interface{}{"collection": "History for Item-" + params["id"]})
}
