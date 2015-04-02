package main

import "github.com/pivotalservices/pezsearch"

func main() {
	s := pezsearch.NewServer()
	s.Run()
}
