package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

//TypeController - controller for searching type information
type TypeController struct {
}

// @Title getTypes
// @Description Returns a list of Types
// @Accept  json
// @Param   token   header  string  true   "Access Token"
// @Success 200 {array}  string
// @Resource /types
// @Router /v1/types [get]
func (c *TypeController) ListTypes(render render.Render) {
	render.JSON(200, map[string]interface{}{"collection": "List of Types"})
}

func (c *TypeController) GetType(params martini.Params, render render.Render) {
	render.JSON(200, map[string]interface{}{"type": "Details of Type-" + params["id"]})
}

func (c *TypeController) ListTypeItems(params martini.Params, render render.Render) {
	render.JSON(200, map[string]interface{}{"collection": "List of Items for Type-" + params["id"]})
}
