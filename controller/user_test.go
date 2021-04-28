package controller

import (
	"testing"

	"aihealth/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestGetAccount(t *testing.T) {
	var err error
	model.MongoSession, err = mgo.Dial("aiwac.net:27017")
	if err != nil {
		t.Errorf("MongoDB connection is not setting")
	}
	var data []interface{}
	collection := model.MongoSession.DB("AIHealth").C("users")
	if err := collection.Find(bson.M{}).All(&data); err == nil {
	} else {
		t.Errorf("GetAccount has error")
	}
}
