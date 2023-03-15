package main

import (
	"fmt"
	"log"
	"net/http"
)

// struct to hold server info
type APIServer struct {
	listenAddr string
	// db interface here
	db *PostgresDB
}

// function to create new API Server
func NewAPIServer(listenAddr string, db *PostgresDB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

// function signature to be used in this app
type apifunc func(http.ResponseWriter, *http.Request) error

// decorate our apifunc into a handlerfunc
func makeHTTPHandler(f apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *APIServer) Run() {

	mux := http.NewServeMux()

	mux.Handle("/", makeHTTPHandler(s.handleHome))

	log.Println("starting server on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, mux)

}

func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) error {

	fmt.Println(JSONResponse{Message: "hello world"})

	return nil

}
