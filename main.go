package main

import (
	"aihealth/model"
	"aihealth/router"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDB() (context.Context, error) {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://aihealth:YWaMxPZh7mGM5k7Q@cluster0.xkwrz.mongodb.net/AIHealth?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	model.MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return ctx, err
}

// @title AIHealth API
// @version 1.0
// @description This is a AIHealth server.
// @host localhost:10086
// @BasePath /
func main() {
	log.Println("AIHealth start...")
	log.Println("Connecting MongoDB")
	ctx, err := connectMongoDB()
	if err != nil {
		log.Fatal("Connect MongoDB Failed")
		log.Println(err)
	} else {
		log.Println("Connected MongoDB Success")
	}
	defer func() {
		if err = model.MongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	r := router.SetupRouter()
	r.Run(":10086")
}
