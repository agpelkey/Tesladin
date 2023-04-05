package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() (*MongoInstace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//dbPasswd := os.Getenv("$MONGO_PASSWD")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://apelkey:CombustionSorc2@cluster0.zpbcywr.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	/*
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	*/

	db := client.Database("FileStorage").Collection("files")

	mg := MongoInstace{
		Client: client,
		Db:     db.Database(),
	}

	return &mg, nil

}
