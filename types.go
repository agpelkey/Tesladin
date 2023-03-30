package main

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostgresDB struct {
	db *sql.DB //database connection pool
}

type JSONResponse struct {
	Message string `json:"message"`
}

type MongoInstace struct {
	Client *mongo.Client
	Db     *mongo.Database
}
