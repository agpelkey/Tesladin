package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {

	/*
		dbconn, err := NewPostgresDB()
		if err != nil {
			log.Fatal(err)
		}

		if err := dbconn.Init(); err != nil {
			log.Fatal(err)
		}
	*/
	dbConn, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8080", dbConn)

	server.Run()

}
