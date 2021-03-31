package dbs

import (
	"log"

	"gopkg.in/mgo.v2"
)

func InsertMongo(mongo *mgo.Session, DBname string, collection string, data interface{}) error {
	client := mongo.DB(DBname).C(collection)
	log.Println("Insert data: ", data)

	// cErr := client.Insert(model.User{
	// 	User_ID: "1234",
	// 	Name:    "name",
	// 	Birth:   "birth",
	// 	Gender:  "gender",
	// 	ABO:     "abo",
	// 	Phone:   "phone",
	// 	Rh:      false,
	// 	Height:  165,
	// 	Weight:  65,
	// })

	cErr := client.Insert(data)

	return cErr
}
