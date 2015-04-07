package main

import pez "github.com/pivotalservices/pezsearch/service"

func main() {
	s := pez.NewServer()
	s.Run()
}
