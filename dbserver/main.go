package main

import (
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"lxh027.com/xml-dbms/dbserver/grpc_server"
)

func main()  {
	runtime.LoadDataFromXML()
	grpc_server.StartRpcServer()
}