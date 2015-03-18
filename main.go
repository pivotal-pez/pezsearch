// @APIVersion 1.0.0
// @APITitle Pez Inventory Search
// @APIDescription Search Pez inventory with these APIs
// @SubApi Type Search [/types]
// @SubApi Item Search [/items]
package main

import "github.com/pivotalservices/pezinventory/pezinventory"

func main() {
	server := pezinventory.NewServer()
	server.Run()
}
