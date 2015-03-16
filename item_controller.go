package pezinventory

import "github.com/go-martini/martini"

//ItemController - controller for searching item information
type ItemController struct {
}

func (c *ItemController) ListItems() string {
	return "List of Items"
}

func (c *ItemController) GetItem(params martini.Params) string {
	return "Details of Item-" + params["id"]
}

func (c *ItemController) GetItemHistory(params martini.Params) string {
	return "History for Item-" + params["id"]
}
