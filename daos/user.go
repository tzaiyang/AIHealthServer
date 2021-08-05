package daos

import (
	"aihealth/common"
	"aihealth/databases"
	"aihealth/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User manages User CRUD
type User struct {
}

func (u *User) Find(filter interface{}) (interface{}, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColUsers)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	err = cursor.All(context.TODO(), &results)

	return results, err
}

func (u *User) InsertOne(user models.User) interface{} {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColUsers)
	log.Println("Insert data: ", user)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return res.InsertedID
}

func (u *User) UpdateByID(user_id string, user models.User) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColUsers)
	log.Println("Update data: ", user)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"name": user_id}

	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, user, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		log.Println("matched and replaced an existing document")
		return
	}
	if result.UpsertedCount != 0 {
		log.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
}

func (u *User) UserDeleteByID(user_id string) int64 {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColUsers)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"user_id": user_id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return result.DeletedCount
}
