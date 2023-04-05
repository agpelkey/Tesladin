package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

// struct to hold server info
type APIServer struct {
	listenAddr string
	// db interface here
	db *mongo.Database
}

// function to create new API Server
func NewAPIServer(listenAddr string, mg *MongoInstace) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         mg.Db,
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

	mux.Handle("/", makeHTTPHandler(s.handleFile))

	log.Println("starting server on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, mux)

}

func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) error {

	fmt.Println(JSONResponse{Message: "hello world"})

	return nil

}

func (s *APIServer) handleFile(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		// logic for sending file post request.
		filePayload := File{}
		payload := json.NewDecoder(r.Body).Decode(&filePayload)

		coll := s.db.Client().Database("FileServer").Collection("files")

		result, err := coll.InsertOne(context.TODO(), payload)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
	return nil
}

func (s *APIServer) RetrieveFile(w http.ResponseWriter, r *http.Request) (*File, error) {
	if r.Method == "GET" {
		id := r.URL.Query()

		file := File{}

		coll := s.db.Client().Database("FileServer")

		if err := coll.Collection("FileServer").FindOne(context.TODO(), id).Decode(&file); err != nil {
			log.Fatal(err)
		}

		return &file, nil
	}

	return nil, fmt.Errorf("Method %s not allowed", r.Method)
}
