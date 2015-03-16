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
