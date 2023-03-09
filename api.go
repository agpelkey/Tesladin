package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// struct to hold server info
type APIServer struct {
	listenAddr string
	// db interface here
}

// function to create new API Server
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
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
	io.WriteString(w, "Welcome to the home page")
	return nil
}
