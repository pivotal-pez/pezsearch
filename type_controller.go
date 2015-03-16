package pezinventory

import "github.com/go-martini/martini"

//TypeController - controller for searching type information
type TypeController struct {
}

func (c *TypeController) ListTypes() string {
	return "List of Types"
}

func (c *TypeController) GetType(params martini.Params) string {
	return "Details of Type-" + params["id"]
}

func (c *TypeController) ListTypeItems(params martini.Params) string {
	return "List of Items for Type-" + params["id"]
}
