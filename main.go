package main

import (
	_ "aihealth/docs" // docs is generated by Swag CLI, you have to import it.
	"aihealth/model"
	"aihealth/router"
	"crypto/tls"
	"log"
	"net"

	//GORM allows customize the MySQL driver with the DriverName option, for example:
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/mgo.v2"
)

func ConnectMongoDB() (*mgo.Session, error) {
	dialInfo := mgo.DialInfo{
		Addrs: []string{
			"cluster0-shard-00-00.xkwrz.mongodb.net:27017",
			"cluster0-shard-00-01.xkwrz.mongodb.net:27017",
			"cluster0-shard-00-02.xkwrz.mongodb.net:27017"},
		Username: "aihealth",
		Password: "YWaMxPZh7mGM5k7Q",
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig) // add TLS config
		return conn, err
	}
	var err error
	// model.MongoSession, err = mgo.DialWithInfo(&dialInfo)
	model.MongoSession, err = mgo.Dial("127.0.0.1")
	// model.MongoSession, err = mgo.Dial("aiwac.net:27017")
	return model.MongoSession, err
}

// @title AIHealth API
// @version 1.0
// @description This is a AIHealth server.
// @host localhost:10086
// @BasePath /
func main() {
	log.Println("AIHealth start...")
	log.Println("Connecting MongoDB")
	_, err := ConnectMongoDB()
	if err != nil {
		log.Fatal("Connect MongoDB Failed")
		log.Println(err)
	} else {
		log.Println("Connected MongoDB Success")
	}
	defer model.MongoSession.Close()

	log.Println("Connecting MySQL")

	// model.MySQLPool, err = sql.Open("mysql", dbw.DSN)

	dsn := "root:966841@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	model.MySQLPool, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("Unable to use data source name", err)
		panic(err)
	}
	// defer model.MySQLPool.Close()

	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	model.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pool := goredis.NewPool(model.RedisClient)
	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	model.RediSync = redsync.New(pool)

	r := router.SetupRouter()
	r.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":10086")
}
