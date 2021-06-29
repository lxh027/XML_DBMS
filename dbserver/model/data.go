package model

import "encoding/xml"

type AllData struct {
	XMLName	xml.Name 	`xml:"rows"`
	Rows 	[]Row 		`xml:"row"`
}

type Row struct {
	Data 	[]string 	`xml:"data"`
}
