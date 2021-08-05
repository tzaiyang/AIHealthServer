package common

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Mongo struct {
		URI    string `yaml:"uri"`
		DbName string `yaml:"dbname"`
	}
}

// Config shares the global configuration
var (
	Config *Configuration
)

// COLLECTIONs of the database table
const (
	ColMtrs    = "mtrs"
	ColMedical = "medicals"
	ColUsers   = "users"
)

func LoadConfig() error {
	Config = new(Configuration)

	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v when to struct", err)
	}
	return nil
}
