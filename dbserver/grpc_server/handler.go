package grpc_server

import (
	"log"
	"lxh027.com/xml-dbms/config"
	"lxh027.com/xml-dbms/dbserver/data/handler"
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"lxh027.com/xml-dbms/dbserver/parser/parsed_data"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
	"lxh027.com/xml-dbms/proto"
)

func handleDatabase(parsedBasicData *parsed_data.ParsedBasicData) (*proto.SqlResult, error) {
	var databaseHandler handler.DataBaseHandler
	databaseHandler.Name = parsedBasicData.Name
	databaseHandler.Operation = parsedBasicData.Operation
	if parsedBasicData.Operation == tokenizer.Create {
		databaseHandler.Location = config.DbConfig.DatabasePath+parsedBasicData.Name+".xml"
	} else if parsedBasicData.Operation == tokenizer.Drop {
		index, ok := runtime.GetDatabaseInfoIndex(parsedBasicData.Name)
		if !ok {
			return &proto.SqlResult{
				Status: proto.SqlResult_Syntax_Error,
				Message: "database不存在",
				MetaData: []string{"message"},
				Data: []*proto.SqlResult_DataRow{{DataCell: []string{"database不存在"}}},
			}, nil
		}
		databaseHandler.Location = runtime.Server.DataBases[index].Location
	}
	if err := databaseHandler.ExecSql(); err != nil {
		log.Printf("error: %v\n", err.Error())
		return &proto.SqlResult{
			Status: proto.SqlResult_Sql_Error,
			Message: "运行错误",
			MetaData: []string{"message"},
			Data: []*proto.SqlResult_DataRow{{DataCell: []string{err.Error()}}},
		}, nil
	}
	return &proto.SqlResult{
		Status: proto.SqlResult_OK,
		Message: "运行成功",
		MetaData: []string{"message"},
		Data: []*proto.SqlResult_DataRow{{DataCell: []string{"ok"}}},
	}, nil
}

func handleTableDrop(parsedBasicData *parsed_data.ParsedBasicData) (*proto.SqlResult, error) {
	var tableHandler handler.TableHandler
	tableHandler.Name = parsedBasicData.Name
	tableHandler.Operation = parsedBasicData.Operation
	tableHandler.Location = config.DbConfig.TablePath + runtime.UsedDatabase+"."+parsedBasicData.Name+".xml"
	if err := tableHandler.ExecSql(); err != nil {
		return &proto.SqlResult{
			Status: proto.SqlResult_Sql_Error,
			Message: "运行错误",
			MetaData: []string{"message"},
			Data: []*proto.SqlResult_DataRow{{DataCell: []string{err.Error()}}},
		}, nil
	}
	return &proto.SqlResult{
		Status: proto.SqlResult_OK,
		Message: "运行成功",
		MetaData: []string{"message"},
		Data: []*proto.SqlResult_DataRow{{DataCell: []string{"ok"}}},
	}, nil
}

func handlerTableCreate(table *parsed_data.ParsedCreateTable) (*proto.SqlResult, error)  {
	var tableHandler handler.TableHandler
	tableHandler.Name = table.Name
	tableHandler.Operation = table.Operation
	tableHandler.Location = config.DbConfig.TablePath + runtime.UsedDatabase+"."+table.Name+".xml"
	tableHandler.Columns = table.Columns
	if err := tableHandler.ExecSql(); err != nil {
		return &proto.SqlResult{
			Status: proto.SqlResult_Sql_Error,
			Message: "运行错误",
			MetaData: []string{"message"},
			Data: []*proto.SqlResult_DataRow{{DataCell: []string{err.Error()}}},
		}, nil
	}
	return &proto.SqlResult{
		Status: proto.SqlResult_OK,
		Message: "运行成功",
		MetaData: []string{"message"},
		Data: []*proto.SqlResult_DataRow{{DataCell: []string{"ok"}}},
	}, nil
}

func handlerUse(table *parsed_data.ParsedBasicData) (*proto.SqlResult, error)  {
	var databaseHandler handler.DataBaseHandler
	databaseHandler.Name = table.Name
	databaseHandler.Operation = table.Operation
	if err := databaseHandler.ExecSql(); err != nil {
		return &proto.SqlResult{
			Status: proto.SqlResult_Sql_Error,
			Message: "运行错误",
			MetaData: []string{"message"},
			Data: []*proto.SqlResult_DataRow{{DataCell: []string{err.Error()}}},
		}, nil
	}
	return &proto.SqlResult{
		Status: proto.SqlResult_OK,
		Message: "运行成功",
		MetaData: []string{"message"},
		Data: []*proto.SqlResult_DataRow{{DataCell: []string{"ok"}}},
	}, nil
}