package handler

import (
	"encoding/xml"
	"errors"
	"lxh027.com/xml-dbms/dbserver/data/model"
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
	"os"
)

type DataBaseHandler struct {
	Location 	string
	Name 		string
	Operation 	tokenizer.TokenType
}

func (handler *DataBaseHandler) ExecSql() error {
	if err := handler.SaveToRuntime(); err != nil {
		return nil
	}
	return handler.SaveToXMLFile()
}

func (handler *DataBaseHandler) SaveToRuntime() error {
	switch handler.Operation {
	case tokenizer.Create:
		{
			if _, ok := runtime.Databases[handler.Name]; ok {
				return errors.New("database existed")
			}
			runtime.Server.DataBases = append(runtime.Server.DataBases,
				model.DataBaseInfo{Name: handler.Name, Location: handler.Location})
			var database model.DataBase
			database.Name = handler.Name
			database.Tables = make([]model.Table, 0)
			database.Views = make([]model.View, 0)
			runtime.Databases[handler.Name] = &database
			return nil
		}
	case tokenizer.Drop:
		{
			if _, ok := runtime.Databases[handler.Name]; !ok {
				return errors.New("database non-existed")
			}
			var index int
			for index = 0; index < len(runtime.Server.DataBases); index++ {
				if runtime.Server.DataBases[index].Name == handler.Name {
					break
				}
			}
			runtime.Server.DataBases = append(runtime.Server.DataBases[:index-1],
				runtime.Server.DataBases[index+1:]...)
			delete(runtime.Databases, handler.Name)
			return nil
		}
	default:
		return errors.New("unknown error")
	}
}

func (handler *DataBaseHandler) SaveToXMLFile() error {
	switch handler.Operation {
	case tokenizer.Create:
		{
			databaseXmlFile, err := os.Create(handler.Location)
			if databaseXmlFile != nil {
				defer databaseXmlFile.Close()
			}
			if err != nil {
				return err
			}
			encoder := xml.NewEncoder(databaseXmlFile)
			err = encoder.Encode(runtime.Databases[handler.Name])
			if err != nil {
				return err
			}
			return SaveRootXmlFile()
		}
	case tokenizer.Drop:
		{
			if err := os.Remove(handler.Location); err != nil {
				return nil
			}
			return SaveRootXmlFile()
		}
	default:
		return errors.New("unknown error")
	}
}

