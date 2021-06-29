package model

import "encoding/xml"

type Server struct {
	XMLName 	xml.Name        `xml:"server"`
	Password 	string          `xml:"password"`
	DataBases 	[]dataBaseInfo 	`xml:"database"`
}

type dataBaseInfo struct {
	Location 	string 		`xml:"location"`
	Name 		string 		`xml:"name"`
}
