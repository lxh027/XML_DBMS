package model

import "encoding/xml"

type AllData struct {
	XMLName	xml.Name `xml:"rows"`
	Rows 	[]row       `xml:"row"`
}

type row struct {
	Data 	[]string 	`xml:"data"`
}
