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
func NewAPIServer(listenAddr string, db *MongoInstace) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db.Db,
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

/*
func (s *APIServer) FileUpload(file, filename string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	bucket, err := gridfs.NewBucket(
		conn.Client.Database("files"),
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	payload, err := bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer payload.Close()

	fileSize, err := payload.Write(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("Write file to DB was successfull. File size: %d M\n", fileSize)
}
*/
