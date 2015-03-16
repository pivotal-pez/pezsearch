// @APIVersion 1.0.0
// @APITitle Pez Inventory Search
// @APIDescription Search Pez inventory with these APIs
// @SubApi Type Search [/types]
// @SubApi Item Search [/items]
package main

import (
	"github.com/go-martini/martini"
	"github.com/pivotalservices/pezinventory"
)

func main() {
	m := martini.Classic()
	pezinventory.InitRoutes(m)
	m.Run()
}
