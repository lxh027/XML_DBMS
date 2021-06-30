package handler

import (
	"encoding/xml"
	"lxh027.com/xml-dbms/config"
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"os"
)

type Handler interface {
	SaveToRuntime() error
	SaveToXMLFile() error
	ExecSql() error
}

func SaveRootXmlFile() error {
	systemXmlFile, err := os.OpenFile(config.DbConfig.Root, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if systemXmlFile != nil {
		defer systemXmlFile.Close()
	}
	if err != nil {
		return err
	}
	encoder := xml.NewEncoder(systemXmlFile)
	err = encoder.Encode(runtime.Server)
	if err != nil {
		return err
	}
	return nil
}