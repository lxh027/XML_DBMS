package model

import "encoding/xml"

type DataBase struct {
	XMLName xml.Name 	`xml:"database"`
	Name 	string 		`xml:"name"`
	Tables 	[]Table 	`xml:"table"`
}

type Table struct {
	XMLName		xml.Name 	`xml:"table"`
	Location	string 		`xml:"location"`
	Columns 	[]Column 	`xml:"column"`
	Indexes 	[]Column 	`xml:"index"`
}

type Column struct {
	Name 	string 	`xml:"name"`
	Type 	string 	`xml:"type"`
}