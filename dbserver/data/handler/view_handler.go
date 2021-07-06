package handler

import (
	"errors"
	"lxh027.com/xml-dbms/dbserver/data/model"
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
)

type ViewHandler struct {
	Location 	string
	Name 		string
	Operation 	tokenizer.TokenType
	Sql 		string
}

func (handler *ViewHandler) ExecSql() error {
	if err := handler.SaveToRuntime(); err != nil {
		return err
	}
	return handler.SaveToXMLFile()
}

func (handler *ViewHandler) SaveToRuntime() error {
	index, ok := runtime.Databases[runtime.UsedDatabase].GetViewIndex(handler.Name)
	switch handler.Operation {
	case tokenizer.Create:
		{
			if ok {
				return errors.New("view existed")
			}
			runtime.Databases[runtime.UsedDatabase].Views = append(runtime.Databases[runtime.UsedDatabase].Views,
				model.View{Name: handler.Name, Sql: handler.Sql})
			return nil
		}
	case tokenizer.Drop:
		{
			if !ok {
				return errors.New("database non-existed")
			}
			if index == 0 {
				if len(runtime.Databases[runtime.UsedDatabase].Views) == 1 {
					runtime.Databases[runtime.UsedDatabase].Views = make([]model.View, 0)
				} else {
					runtime.Databases[runtime.UsedDatabase].Views = runtime.Databases[runtime.UsedDatabase].Views[index+1:]
				}
			} else if index == len(runtime.Databases[runtime.UsedDatabase].Views)-1 {
				runtime.Databases[runtime.UsedDatabase].Views = runtime.Databases[runtime.UsedDatabase].Views[:index]
			} else {
				runtime.Databases[runtime.UsedDatabase].Views =
					append(runtime.Databases[runtime.UsedDatabase].Views[:index],
						runtime.Databases[runtime.UsedDatabase].Views[index+1:]...)
			}
			return nil
		}
	default:
		return errors.New("unknown error")
	}
}

func (handler *ViewHandler) SaveToXMLFile() error {
	databaseInfoIndex, ok := runtime.GetDatabaseInfoIndex(runtime.UsedDatabase)
	if !ok {
		return errors.New("database index error")
	}
	switch handler.Operation {
	case tokenizer.Create:
		{
			return runtime.Databases[runtime.UsedDatabase].SaveToXmlFile(runtime.Server.DataBases[databaseInfoIndex].Location)
		}
	case tokenizer.Drop:
		{
			return runtime.Databases[runtime.UsedDatabase].SaveToXmlFile(runtime.Server.DataBases[databaseInfoIndex].Location)
		}
	default:
		return errors.New("unknown error")
	}
}
