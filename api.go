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
	mux.Handle("/tesladin/{id}", makeHTTPHandler(s.handleRetrieveFile))

	log.Println("starting server on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, mux)

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

		response := fmt.Sprintf("http://localhost:8080/tesladin/%s", result.InsertedID)

		return WriteJSON(w, http.StatusOK, response)

	}
	return nil
}

func (s *APIServer) handleRetrieveFile(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id := r.URL.Query()

		file := File{}

		coll := s.db.Client().Database("FileServer")

		if err := coll.Collection("FileServer").FindOne(context.TODO(), id).Decode(&file); err != nil {
			log.Fatal(err)
		}

		return WriteJSON(w, http.StatusOK, file)

	}

	return fmt.Errorf("Error: method %s not allowed", r.Method)

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
