package main

import "github.com/gophersiesta/gophersiesta/server"

func main() {

	s := server.StartServer()

	defer s.Storage.Close()
}
