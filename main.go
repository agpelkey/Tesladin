package main

import "database/sql"

type JSONResponse struct {
	Message string `json:"message"`
}

func main() {

	db, err := sql.Open("postgres", "fileserver")

	server := NewAPIServer(":8080")

	server.Run()
	
}
