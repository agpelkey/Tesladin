package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {

	/*
		dbconn, err := NewPostgresDB()
		if err != nil {
			log.Fatal(err)
		}

		if err := dbconn.Init(); err != nil {
			log.Fatal(err)
		}
	*/
	dbConn, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	/*
		// create client to connect to deployment that is given by URI
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://apelkey:CombustionSorc2@cluster0.zpbcywr.mongodb.net/?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}

		// create context
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		// Connect initliazes the client
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		// send a ping to verify
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
	*/

	server := NewAPIServer(":8080", dbConn)

	server.Run()

}
