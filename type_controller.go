package pezinventory

import "github.com/go-martini/martini"

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
func (c *TypeController) ListTypes() string {
	return "List of Types"
}

func (c *TypeController) GetType(params martini.Params) string {
	return "Details of Type-" + params["id"]
}

func (c *TypeController) ListTypeItems(params martini.Params) string {
	return "List of Items for Type-" + params["id"]
}
