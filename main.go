package main

import (
	"crud/src/database"
	"crud/src/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()
	s.Run()
}
