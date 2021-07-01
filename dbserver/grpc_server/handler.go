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

func handleTableCreate(table *parsed_data.ParsedCreateTable) (*proto.SqlResult, error)  {
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

func handleViewDrop(parsedBasicData *parsed_data.ParsedBasicData) (*proto.SqlResult, error) {
	var viewHandler handler.ViewHandler
	viewHandler.Name = parsedBasicData.Name
	viewHandler.Operation = parsedBasicData.Operation
	if err := viewHandler.ExecSql(); err != nil {
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

func handleViewCreate(table *parsed_data.ParsedCreateView) (*proto.SqlResult, error)  {
	var viewHandler handler.ViewHandler
	viewHandler.Name = table.Name
	viewHandler.Operation = table.Operation
	viewHandler.Sql = table.Sql
	if err := viewHandler.ExecSql(); err != nil {
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

func handleShow(table *parsed_data.ParsedBasicData) (*proto.SqlResult, error)  {
	switch table.Name {
	case "databases":
		{
			databases := make([]*proto.SqlResult_DataRow, 0)
			for _, databaseInfo := range runtime.Server.DataBases {
				databaseName := databaseInfo.Name
				if databaseName == runtime.UsedDatabase {
					databaseName = "* "+databaseName
				}
				dataRow := proto.SqlResult_DataRow{DataCell: []string{databaseName}}
				databases = append(databases, &dataRow)
			}
			return &proto.SqlResult{
				Status: proto.SqlResult_OK,
				Message: "运行成功",
				MetaData: []string{"database"},
				Data: databases,
			}, nil
		}
	case "tables":
		{
			if runtime.UsedDatabase == "" {
				return &proto.SqlResult{
					Status: proto.SqlResult_Sql_Error,
					Message: "运行错误",
					MetaData: []string{"message"},
					Data: []*proto.SqlResult_DataRow{{DataCell: []string{"database unselected"}}},
				}, nil
			}
			tables := make([]*proto.SqlResult_DataRow, 0)
			for _, tableInfo := range runtime.Databases[runtime.UsedDatabase].Tables {
				dataRow := proto.SqlResult_DataRow{DataCell: []string{tableInfo.Name}}
				tables = append(tables, &dataRow)
			}
			return &proto.SqlResult{
				Status: proto.SqlResult_OK,
				Message: "获取成功",
				MetaData: []string{"table"},
				Data: tables,
			}, nil
	}
	case "views":
		{
			if runtime.UsedDatabase == "" {
				return &proto.SqlResult{
					Status: proto.SqlResult_Sql_Error,
					Message: "运行错误",
					MetaData: []string{"message"},
					Data: []*proto.SqlResult_DataRow{{DataCell: []string{"database unselected"}}},
				}, nil
			}
			views := make([]*proto.SqlResult_DataRow, 0)
			for _, viewInfo := range runtime.Databases[runtime.UsedDatabase].Views {
				dataRow := proto.SqlResult_DataRow{DataCell: []string{viewInfo.Name, viewInfo.Sql}}
				views = append(views, &dataRow)
			}
			return &proto.SqlResult{
				Status: proto.SqlResult_OK,
				Message: "获取成功",
				MetaData: []string{"view", "sql"},
				Data: views,
			}, nil
		}
	}
	return &proto.SqlResult{
		Status: proto.SqlResult_Sql_Error,
		Message: "运行错误",
		MetaData: []string{"message"},
		Data: []*proto.SqlResult_DataRow{{DataCell: []string{"show的对象不存在"}}},
	}, nil
}