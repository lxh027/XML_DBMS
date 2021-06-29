package model

import "encoding/xml"

type DataBase struct {
	XMLName xml.Name  `xml:"database"`
	Name 	string    `xml:"name"`
	Tables 	[]Table `xml:"table"`
}

type Table struct {
	XMLName		xml.Name `xml:"table"`
	Name 		string 		`xml:"name"`
	Location	string      `xml:"location"`
	Columns 	[]column     `xml:"column"`
	Indexes 	[]column     `xml:"index"`
}

type column struct {
	Name 	string 	`xml:"name"`
	Type 	string 	`xml:"type"`
}