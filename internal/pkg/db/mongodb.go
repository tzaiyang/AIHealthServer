package db

import (
	"aihealth/internal/pkg/cfg"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbClient *mongo.Client
	DbName     string
}

func (db *MongoDB) Init() error {
	db.DbName = cfg.Config.Mongo.DbName
	clientOptions := options.Client().
		ApplyURI(cfg.Config.Mongo.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	db.MgDbClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	return nil
}

// func (db *MongoDB) InsertOne(document interface{}) (interface{}, error) {
// 	collection := db.MgDbClient.Database(db.DbName).Collection(common.ColMedical)
// 	// log.Println("Insert document: ", document)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	res, err := collection.InsertOne(ctx, document)

// 	return res.InsertedID, err
// }

// // filter := bson.M{"user_id": user_id}
// func (db *MongoDB) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
// 	collection := db.MgDbClient.Database(db.DbName).Collection(common.ColUsers)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	return collection.DeleteOne(ctx, filter)
// }

// func (db *MongoDB) Delete(filter interface{}) (int64, error) {
// 	collection := db.MgDbClient.Database(db.DbName).Collection(common.ColMtrs)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	results, err := collection.DeleteOne(ctx, filter)

// 	return results.DeletedCount, err
// }

// func (db *MongoDB) Find(colName string, filter interface{}) (interface{}, error) {
// 	collection := db.MgDbClient.Database(common.Config.Mongo.DbName).Collection(colName)
// 	var results []bson.M

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	cursor, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = cursor.All(context.TODO(), &results)

// 	return results, err
// }

func (db *MongoDB) Close() {
	if db.MgDbClient != nil {
		db.MgDbClient.Disconnect(context.Background())
	}
}
