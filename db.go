package main

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type PostgresRepo interface {
	NewFile(name string) (int64, error)
}

const dbtimeout = time.Second * 3

func (m *PostgresDB) Connection() *sql.DB {
	return m.db
}

func NewPostgresDB() (*PostgresDB, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=fileserver sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresDB{
		db: db,
	}, nil

}

func (m *PostgresDB) CreateTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS files (
		id SERIAL PRIMARY KEY
	);`

	_, err := m.db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (m *PostgresDB) Init() error {
	return m.CreateTable()
}

func (m *PostgresDB) NewFile(name string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	stmt, err := m.db.ExecContext(ctx, "INSERT INTO files VALUES (name = ?)", name)
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}
