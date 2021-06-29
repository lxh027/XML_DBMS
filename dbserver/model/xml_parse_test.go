package model

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"testing"
)

func TestXmlServerParse(t *testing.T)  {
	filePos := "../xml/samples/server.xml"
	xmlData, err := ioutil.ReadFile(filePos)
	if err != nil {
		log.Panicf(err.Error())
	}
	var serverData Server
	err = xml.Unmarshal(xmlData, &serverData)
	if err != nil {
		log.Panicf(err.Error())
	}
	log.Println(serverData)
}

func TestXmlDatabaseParse(t *testing.T)  {
	filePos := "../xml/samples/database.xml"
	xmlData, err := ioutil.ReadFile(filePos)
	if err != nil {
		log.Panicf(err.Error())
	}
	var dataBaseData DataBase
	err = xml.Unmarshal(xmlData, &dataBaseData)
	if err != nil {
		log.Panicf(err.Error())
	}
	log.Println(dataBaseData)
}

func TestXmlDataParse(t *testing.T)  {
	filePos := "../xml/samples/data.xml"
	xmlData, err := ioutil.ReadFile(filePos)
	if err != nil {
		log.Panicf(err.Error())
	}
	var data AllData
	err = xml.Unmarshal(xmlData, &data)
	if err != nil {
		log.Panicf(err.Error())
	}
	log.Println(data)
}