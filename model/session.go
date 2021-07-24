package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DbWorker struct {
	//mysql data source name
	DSN string
}

var MongoClient *mongo.Client
