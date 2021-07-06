package handler

import (
	"encoding/xml"
	"errors"
	"lxh027.com/xml-dbms/dbserver/data/model"
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
	"os"
)

type TableHandler struct {
	Location 	string
	Name 		string
	Operation 	tokenizer.TokenType
	Columns 	[]model.Column
}

func (handler *TableHandler) ExecSql() error {
	if err := handler.SaveToRuntime(); err != nil {
		return err
	}
	return handler.SaveToXMLFile()
}

func (handler *TableHandler) SaveToRuntime() error {
	if runtime.UsedDatabase == "" {
		return errors.New("database not selected")
	}
	tableName := runtime.UsedDatabase+"."+handler.Name
	switch handler.Operation {
	case tokenizer.Create:
		{
			if _, ok := runtime.Tables[tableName]; ok {
				return errors.New("table existed")
			}
			runtime.Databases[runtime.UsedDatabase].Tables =
				append(runtime.Databases[runtime.UsedDatabase].Tables,
					model.Table{Name: handler.Name, Location: handler.Location, Columns: handler.Columns})
			var table model.AllData
			table.Name = handler.Name
			table.Rows = make([]model.Row, 0)
			runtime.Tables[tableName] = &table
			return nil
		}
	case tokenizer.Drop:
		{
			if _, ok := runtime.Tables[tableName]; !ok {
				return errors.New("table non-existed")
			}
			index, ok := runtime.Databases[runtime.UsedDatabase].GetTableIndex(handler.Name)
			if !ok {
				return errors.New("table non-existed")
			}
			if index == 0 {
				if len(runtime.Databases[runtime.UsedDatabase].Tables) == 1 {
					runtime.Databases[runtime.UsedDatabase].Tables = make([]model.Table, 0)
				} else {
					runtime.Databases[runtime.UsedDatabase].Tables = runtime.Databases[runtime.UsedDatabase].Tables[index+1:]
				}
			} else if index == len(runtime.Databases[runtime.UsedDatabase].Tables)-1 {
				runtime.Databases[runtime.UsedDatabase].Tables = runtime.Databases[runtime.UsedDatabase].Tables[:index]
			} else {
				runtime.Databases[runtime.UsedDatabase].Tables =
					append(runtime.Databases[runtime.UsedDatabase].Tables[:index],
						runtime.Databases[runtime.UsedDatabase].Tables[index+1:]...)
			}
			delete(runtime.Tables, tableName)
			return nil
		}
	default:
		return errors.New("unknown error")
	}
}

func (handler *TableHandler) SaveToXMLFile() error {
	tableName := runtime.UsedDatabase+"."+handler.Name
	databaseInfoIndex, ok := runtime.GetDatabaseInfoIndex(runtime.UsedDatabase)
	if !ok {
		return errors.New("database index error")
	}
	switch handler.Operation {
	case tokenizer.Create:
		{
			tableXmlFile, err := os.Create(handler.Location)
			if tableXmlFile != nil {
				defer tableXmlFile.Close()
			}
			if err != nil {
				return err
			}
			encoder := xml.NewEncoder(tableXmlFile)
			err = encoder.Encode(runtime.Tables[tableName])
			if err != nil {
				return err
			}
			return runtime.Databases[runtime.UsedDatabase].SaveToXmlFile(runtime.Server.DataBases[databaseInfoIndex].Location)
		}
	case tokenizer.Drop:
		{
			if err := os.Remove(handler.Location); err != nil {
				return nil
			}
			return runtime.Databases[runtime.UsedDatabase].SaveToXmlFile(runtime.Server.DataBases[databaseInfoIndex].Location)
		}
	default:
		return errors.New("unknown error")
	}
}