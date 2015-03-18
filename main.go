// @APIVersion 1.0.0
// @APITitle Pez Inventory Search
// @APIDescription Search Pez inventory with these APIs
// @SubApi Type Search [/types]
// @SubApi Item Search [/items]
package main

import "github.com/pivotalservices/pezinventory/server"

func main() {
	s := server.NewServer()
	s.Run()
}
