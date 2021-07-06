package handler

import (
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
		return err
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
			index, ok := runtime.GetDatabaseInfoIndex(handler.Name)
			if !ok {
				return errors.New("database non-existed")
			}
			if index == 0 {
				if len(runtime.Server.DataBases) == 1 {
					runtime.Server.DataBases = make([]model.DataBaseInfo, 0)
				} else {
					runtime.Server.DataBases = runtime.Server.DataBases[index+1:]
				}
			} else if index == len(runtime.Server.DataBases)-1 {
				runtime.Server.DataBases = runtime.Server.DataBases[:index]
			} else {
				runtime.Server.DataBases =
					append(runtime.Server.DataBases[:index],
						runtime.Server.DataBases[index+1:]...)
			}
			delete(runtime.Databases, handler.Name)
			return nil
		}
	case tokenizer.Use:
		{
			if _, ok := runtime.Databases[handler.Name]; !ok {
				return errors.New("database non-existed")
			}
			if _, ok := runtime.GetDatabaseInfoIndex(handler.Name);!ok {
				return errors.New("database non-existed")
			}
			runtime.UsedDatabase = handler.Name
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
			err := runtime.Databases[handler.Name].SaveToXmlFile(handler.Location)
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
	case tokenizer.Use:
		{
			return nil
		}
	default:
		return errors.New("unknown error")
	}
}

