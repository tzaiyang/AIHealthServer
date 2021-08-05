package databases

import (
	"aihealth/common"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbClient   *mongo.Client
	Databasename string
}

func (db *MongoDB) Init() error {
	db.Databasename = common.Config.Mongo.DbName
	clientOptions := options.Client().
		ApplyURI(common.Config.Mongo.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	db.MgDbClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (db *MongoDB) Close() {
	if db.MgDbClient != nil {
		db.MgDbClient.Disconnect(context.Background())
	}
}
