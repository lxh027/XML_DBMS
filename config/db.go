package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var DbConfig struct {
	Root 			string 	`yml:"root"`
	Interval 		int 	`yml:"interval"`
	DatabasePath 	string 	`yml:"database_path"`
	TablePath 		string 	`yml:"table_path"`
}

func init() {
	yamlFile, err := ioutil.ReadFile("./config/yml/db.yml")
	if err != nil {
		log.Panicf("reading yaml file error: %v\n", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &DbConfig)

	if err != nil {
		log.Panicf("parsing yaml file error: %v\n", err.Error())
	}
	log.Println("rpc yaml file loaded")
}
