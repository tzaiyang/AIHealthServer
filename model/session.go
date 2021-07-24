package model

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"gopkg.in/mgo.v2"
	"gorm.io/gorm"
)

type DbWorker struct {
	//mysql data source name
	DSN string
}

var MongoSession *mgo.Session
var MySQLPool *gorm.DB
var RedisClient *redis.Client
var RediSync *redsync.Redsync
