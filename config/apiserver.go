package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var ApiServerConfig struct {
	Gin struct{
		Host	string 	`yaml:"host"`
		Port 	string 	`yaml:"port"`
		Mode 	string 	`yaml:"mode"`
	}	`yaml:"gin"`
	Session struct{
		Key 	string 	`yaml:"key"`
		Name 	string 	`yaml:"name"`
		Age 	int 	`yaml:"age"`
		Path 	string 	`yaml:"path"`
	}	`yaml:"session"`
}

func init() {
	yamlFile, err := ioutil.ReadFile("./config/yml/apiserver.yml")
	if err != nil {
		log.Panicf("reading yaml file error: %v\n", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &ApiServerConfig)

	if err != nil {
		log.Panicf("parsing yaml file error: %v\n", err.Error())
	}
	log.Println("rpc yaml file loaded")
}
