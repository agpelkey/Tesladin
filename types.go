package main

import "database/sql"

type PostgresDB struct {
	db *sql.DB //database connection pool
}

type JSONResponse struct {
	Message string `json:"message"`
}

type File struct {
	id int `json:"id"`
}
