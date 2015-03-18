package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotalservices/pezinventory/models"
)

//TypeController - controller for searching type information
type TypeController struct {
}

func (c *TypeController) ListTypes(render render.Render) {
	t := models.ListTypes()
	render.JSON(200, successMessage(t))
}

func (c *TypeController) GetType(params martini.Params, render render.Render) {
	t := models.GetType(params["id"])
	render.JSON(200, successMessage(t))
}

func (c *TypeController) ListTypeItems(params martini.Params, render render.Render) {
	i := models.ListItemsByType(params["id"])
	render.JSON(200, successMessage(i))
}
