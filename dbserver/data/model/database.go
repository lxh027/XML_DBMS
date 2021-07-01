package model

import (
	"encoding/xml"
	"os"
)

type DataBase struct {
	XMLName xml.Name  `xml:"database"`
	Name 	string    `xml:"name"`
	Tables 	[]Table `xml:"table"`
	Views 	[]View	`xml:"view"`
}

type View struct {
	XMLName 	xml.Name 	`xml:"view"`
	Name 		string 		`xml:"name"`
	Sql 		string 		`xml:"sql"`
}

type Table struct {
	XMLName		xml.Name `xml:"table"`
	Name 		string 		`xml:"name"`
	Location	string      `xml:"location"`
	Columns 	[]Column     `xml:"column"`
	Indexes 	[]Column     `xml:"index"`
}

type Column struct {
	Name 	string 	`xml:"name"`
	Type 	string 	`xml:"type"`
}

func (database *DataBase) GetViewIndex(name string) (int, bool)  {
	var index int
	for index = 0; index < len(database.Views); index++ {
		if database.Views[index].Name == name {
			return index, true
		}
	}
	return -1, false
}

func (database *DataBase) GetTableIndex(name string) (int, bool) {
	var index int
	for index = 0; index < len(database.Tables); index++ {
		if database.Tables[index].Name == name {
			return index, true
		}
	}
	return -1, false
}

func (database *DataBase) SaveToXmlFile(location string) error  {
	databaseXmlFile, err := os.Create(location)
	if databaseXmlFile != nil {
		defer databaseXmlFile.Close()
	}
	if err != nil {
		return err
	}
	encoder := xml.NewEncoder(databaseXmlFile)
	return encoder.Encode(database)
}