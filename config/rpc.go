package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var RpcConfig struct {
	Network string	`yaml:"network"`
	Host 	string	`yaml:"host"`
	Port 	string	`yaml:"port"`
}

func init() {
	yamlFile, err := ioutil.ReadFile("./config/yml/rpc.yml")
	if err != nil {
		log.Panicf("reading yaml file error: %v\n", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &RpcConfig)

	if err != nil {
		log.Panicf("parsing yaml file error: %v\n", err.Error())
	}
	log.Println("rpc yaml file loaded")
}