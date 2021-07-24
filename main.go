package main

import (
	"aihealth/model"
	"aihealth/router"
	"context"
	"flag"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

func connectMongoDB(config Config) (context.Context, error) {
	clientOptions := options.Client().
		ApplyURI(config.Mongo.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	model.MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return ctx, err
}

type Config struct {
	Mongo struct {
		URI string `yaml:"uri"`
	}
}

// @title AIHealth API
// @version 1.0
// @description This is a AIHealth server.
// @host localhost:10086
// @BasePath /
func main() {
	config_file := flag.String("config", "config.yaml", "Path of configure file")
	flag.Usage()
	flag.Parse()
	conf := new(Config)

	yamlFile, err := ioutil.ReadFile(*config_file)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v when to struct", err)
	}

	log.Println("AIHealth start...")
	log.Println("Connecting MongoDB")
	ctx, err := connectMongoDB(*conf)
	if err != nil {
		log.Fatal("Connect MongoDB Failed")
		log.Println(err)
	} else {
		log.Println("Connected MongoDB Success")
	}
	defer func() {
		if err = model.MongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	r := router.SetupRouter()
	r.Run(":10086")
}
