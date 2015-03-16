package pezinventory

import "github.com/go-martini/martini"

//ItemController - controller for searching item information
type ItemController struct {
}

// @Title getItems
// @Description returns a list of Items
// @Accept  json
// @Param   token   header  string  true   "Access Token"
// @Success 200 {array}  string
// @Resource /items
// @Router /v1/items [get]
func (c *ItemController) ListItems() string {
	return "List of Items"
}

func (c *ItemController) GetItem(params martini.Params) string {
	return "Details of Item-" + params["id"]
}

func (c *ItemController) GetItemHistory(params martini.Params) string {
	return "History for Item-" + params["id"]
}
