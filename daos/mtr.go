// Medical Treatment Record
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

// manages CRUD
type MedicalTreatmentRecord struct {
}

func (m *MedicalTreatmentRecord) InsertOne(mtr models.MedicalTreatmentRecord) interface{} {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMtrs)
	log.Println("Insert mtr: ", mtr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, mtr)
	if err != nil {
		log.Fatal(err)
	}

	return res.InsertedID
}

func (m *MedicalTreatmentRecord) Find(filter interface{}) ([]bson.M, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMtrs)

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

func (m *MedicalTreatmentRecord) Delete(filter interface{}) (int64, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMtrs)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results, err := collection.DeleteOne(ctx, filter)

	return results.DeletedCount, err
}

func (m *MedicalTreatmentRecord) UpdateByID(id interface{}, mtr models.MedicalTreatmentRecord) (*mongo.UpdateResult, error) {
	collection := databases.Database.MgDbClient.Database(common.Config.Mongo.DbName).Collection(common.ColMtrs)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateResults, err := collection.UpdateByID(ctx, id, bson.D{{"$set", mtr}})

	return updateResults, err
}
