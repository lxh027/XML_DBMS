package model

import "encoding/xml"

type Server struct {
	XMLName 	xml.Name         `xml:"server"`
	Password 	string          `xml:"password"`
	DataBases 	[]DataBaseInfo `xml:"database"`
}

type DataBaseInfo struct {
	Location 	string 		`xml:"location"`
	Name 		string 		`xml:"name"`
}
