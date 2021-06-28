package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"lxh027.com/xml-dbms/api-web/grpc"
	"lxh027.com/xml-dbms/api-web/server"
)
func main() {
	// Test Rpc Conn
	if err := grpc.ExecTestConn(context.Background()); err != nil {
		log.Printf("Fail to open rpc conn: %v\n", err.Error())
	}
	// 启动httpServer
	var httpServer *gin.Engine
	server.Run(httpServer)
}