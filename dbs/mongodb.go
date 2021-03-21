package dbs

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func QueryMongo(mongo *mgo.Session, DBname string, collection string, data interface{}) error {
	// client := mongo.DB(DBname).C(collection)

	// // cErr := client.Find(&data)

	// if cErr != nil {
	// 	fmt.Print("cErr")
	// }
	return nil
}
func InsertMongo(mongo *mgo.Session, DBname string, collection string, data interface{}) error {
	client := mongo.DB(DBname).C(collection)

	cErr := client.Insert(&data)

	if cErr != nil {
		fmt.Print("cErr")
	}
	return cErr
}
