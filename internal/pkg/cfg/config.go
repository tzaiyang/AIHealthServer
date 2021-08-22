package cfg

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Service struct {
		Port string `yaml:"port"`
	}
	Mongo struct {
		URI    string `yaml:"uri"`
		DbName string `yaml:"dbname"`
	}
}

// Config shares the global configuration
var (
	Config *Configuration
)

func LoadConfig() error {
	Config = new(Configuration)

	yamlFile, err := ioutil.ReadFile(os.Getenv("CONFIG_URL"))
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v when to struct", err)
	}
	return nil
}
