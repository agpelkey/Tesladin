package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres dbname=fileserver password=fileserver sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	s := &PostgresDB{db: db}

	server := NewAPIServer(":8080", s)

	server.Run()

}
