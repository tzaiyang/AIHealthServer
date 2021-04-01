package model

import (
	"log"

	"gopkg.in/mgo.v2/bson"
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

func (user User) Insert() error {
	collection := MongoSession.DB("AIHealth").C("users")
	log.Println("Insert data: ", user)
	err := collection.Insert(user)
	return err
}

func (user User) UpdateByID(user_id string) error {
	collection := MongoSession.DB("AIHealth").C("users")
	err := collection.Update(bson.M{"user_id": user_id}, user)
	return err
}

func UserDeleteByID(user_id string) error {
	collection := MongoSession.DB("AIHealth").C("users")
	err := collection.Remove(bson.M{"user_id": user_id})
	return err
}
