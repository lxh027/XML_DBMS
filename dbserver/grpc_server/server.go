package grpc_server

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"lxh027.com/xml-dbms/config"
	"lxh027.com/xml-dbms/dbserver/data/handler"
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"lxh027.com/xml-dbms/dbserver/parser"
	"lxh027.com/xml-dbms/dbserver/parser/parsed_data"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
	"lxh027.com/xml-dbms/proto"
	"net"
	"reflect"
)

type dbRpcServer struct {}

func (server *dbRpcServer) Auth(c context.Context, request *proto.AuthRequest) (*proto.AuthResponse, error)  {
	if request.Password == runtime.Server.Password {
		return &proto.AuthResponse{Message: "验证成功", Status: proto.AuthResponse_OK}, nil
	}
	return &proto.AuthResponse{Message: "验证失败", Status: proto.AuthResponse_Error}, nil
}

func (server *dbRpcServer) SqlExecute(c context.Context, expression *proto.SQLExpression) (*proto.SqlResult, error)  {
	parsed, err := parser.ParseSql(expression.Sql)
	if err != nil {
		return &proto.SqlResult{
			Status: proto.SqlResult_Syntax_Error,
			Message: "解析sql失败",
			MetaData: []string{"message"},
			Data: []*proto.SqlResult_DataRow{{DataCell: []string{err.Error()}}},
		}, nil
	}
	log.Printf("parsed is a type of %v\n", reflect.TypeOf(parsed))
	switch reflect.TypeOf(parsed) {
	case reflect.TypeOf(&parsed_data.ParsedBasicData{}):
		{
			parsedBasicData := parsed.(*parsed_data.ParsedBasicData)
			if parsedBasicData.Target == tokenizer.Database {
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
		}
	case reflect.TypeOf(&parsed_data.ParsedCreateTable{}):
		{

		}
	case reflect.TypeOf(&parsed_data.ParsedCreateView{}):
		{

		}
	}
	panic("")
}

func (server *dbRpcServer) TestConn(c context.Context, ping *proto.Ping) (*proto.Pong, error)  {
	return &proto.Pong{Pong: "hello "+ping.Ping}, nil
}

func StartRpcServer()  {
	lis, err := net.Listen(config.RpcConfig.Network, ":"+config.RpcConfig.Port)
	if err != nil {
		log.Panicf("starting %v port %v error: %v\n",
			config.RpcConfig.Network, config.RpcConfig.Host, err.Error())
	}

	rpcServer := grpc.NewServer()
	proto.RegisterDBServerServer(rpcServer, &dbRpcServer{})

	err = rpcServer.Serve(lis)
	if err != nil {
		log.Panicf("rpc server error: %v\n", err.Error())
	}
	log.Printf("rpc server start at %v on %v", config.RpcConfig.Port, config.RpcConfig.Network)
}



