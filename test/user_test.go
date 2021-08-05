package daos

import (
	"aihealth/common"
	"aihealth/controllers"
	"aihealth/databases"
	"testing"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetAccount(t *testing.T) {
	var err error
	common.LoadConfig()

	log.Println(common.Config.Mongo.DbName)
	if err = databases.Database.Init(); err != nil {
		log.Println("MongoDB connection is not setting")
	}

	defer databases.Database.Close()
	var user controllers.User
	if data, err := user.DAO.Find(bson.M{}); err == nil {
		log.Println(data)
	} else {
		log.Println(err)
	}
}
