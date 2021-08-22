package handler

import (
	"aihealth/internal/pkg/cfg"
	"aihealth/internal/pkg/db"
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetAccount(t *testing.T) {
	var err error
	cfg.LoadConfig()

	log.Println(cfg.Config.Mongo.DbName)
	if err = db.Database.Init(); err != nil {
		log.Println("MongoDB connection is not setting")
	}

	defer db.Database.Close()
	user_name := "tzy"
	user_id := "1234"
	var results []bson.M

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColUsers)
	if cursor, err := collection.Find(context.Background(), bson.M{"name": user_name, "user_id": user_id}); err == nil {
		if err := cursor.All(context.TODO(), &results); err != nil {
			// c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			// c.JSON(http.StatusOK, results)
		}
	} else {
		log.Print(err)
		// c.JSON(http.StatusInternalServerError, err.Error())
	}
}
