package model

import "encoding/xml"

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