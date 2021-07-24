package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	User_ID    string `json:"user_id" example:"18717992222" format:"string"`
	Phone      string `json:"phone" example:"18717992222"`
	Name       string `json:"name" example:"Ryan"`
	Birth      string `json:"birth" example:"2009-08-23"`
	Sex        string `json:"sex" example:"男"`
	ABO        string `json:"abo" example:"B"`
	Rh         bool   `json:"rh" example:"true"`
	Height     int    `json:"height" example:"170"`
	Weight     int    `json:"weight" example:"65"`
	Occupation string `json:"occupation" example:"程序员"`
	Updated    string `json:"updated" example:"2021-03-30 15:59"`
}

func (user User) Find(filter interface{}) ([]interface{}, error) {
	collection := MongoClient.Database("AIHealth").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var results []interface{}
	err = cursor.All(context.TODO(), &results)

	return results, err
}

func (user User) InsertOne() interface{} {
	collection := MongoClient.Database("AIHealth").Collection("users")
	log.Println("Insert data: ", user)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return res.InsertedID
}

func (user User) UpdateByID(user_id string) {
	collection := MongoClient.Database("AIHealth").Collection("users")
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

func UserDeleteByID(user_id string) int64 {
	collection := MongoClient.Database("AIHealth").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"user_id": user_id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return result.DeletedCount
}
