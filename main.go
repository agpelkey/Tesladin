package main

import (
	"log"
)

func main() {

	dbConn, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8080", dbConn)

	server.Run()

}
