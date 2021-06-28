package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"lxh027.com/xml-dbms/config"
	"lxh027.com/xml-dbms/proto"
)

var serverAddr string

func init()  {
	serverAddr = config.RpcConfig.Host+":"+config.RpcConfig.Port
}

func ExecTestConn(c context.Context) error {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return err
	}
	client := proto.NewDBServerClient(conn)

	pong, err := client.TestConn(c, &proto.Ping{Ping: "api web"})
	if err != nil {
		return err
	}
	log.Printf("receive rpc response %v\n", pong.Pong)
	return nil
}

func ExecAuth(c context.Context, password string) (*proto.AuthResponse, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return nil, err
	}
	client := proto.NewDBServerClient(conn)

	return client.Auth(c, &proto.AuthRequest{Password: password})
}

func ExecSql(c context.Context, sql string) (*proto.SqlResult, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return nil, err
	}
	client := proto.NewDBServerClient(conn)

	return client.SqlExecute(c, &proto.SQLExpression{Sql: sql})
}