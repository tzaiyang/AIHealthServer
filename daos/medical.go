package daos

import (
	"aihealth/common"
	"aihealth/databases"
	"aihealth/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Medical struct {
}

func (m *Medical) InsertOne(medical models.Medical) (interface{}, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMedical)
	log.Println("Insert medical: ", medical)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, medical)

	return res.InsertedID, err
}

func (m *Medical) Find(filter interface{}) (interface{}, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMedical)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	err = cursor.All(ctx, &results)

	return results, err
}

func (m *Medical) Delete(filter interface{}) (int64, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMedical)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results, err := collection.DeleteOne(ctx, filter)

	return results.DeletedCount, err
}

func (m *Medical) UpdateByID(id interface{}, medical models.Medical) (*mongo.UpdateResult, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMedical)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateResults, err := collection.UpdateByID(ctx, id, bson.D{{"$set", medical}})

	return updateResults, err
}
