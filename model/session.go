package model

import (
	"gopkg.in/mgo.v2"
	"gorm.io/gorm"
)

type DbWorker struct {
	//mysql data source name
	DSN string
}

var MongoSession *mgo.Session
var MySQLPool *gorm.DB
