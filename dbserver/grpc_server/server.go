package grpc_server

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"lxh027.com/xml-dbms/config"
	"lxh027.com/xml-dbms/proto"
	"net"
)

type dbRpcServer struct {}

func (server *dbRpcServer) Auth(c context.Context, request *proto.AuthRequest) (*proto.AuthResponse, error)  {
	panic("")
}

func (server *dbRpcServer) SqlExecute(c context.Context, expression *proto.SQLExpression) (*proto.SqlResult, error)  {
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



