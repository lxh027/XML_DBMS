package main

import (
	"log"
	"lxh027.com/xml-dbms/dbserver/data/runtime"
	"lxh027.com/xml-dbms/dbserver/grpc_server"
)

func main()  {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover from panic: %v", err)
		}
	}()
	runtime.LoadDataFromXML()
	grpc_server.StartRpcServer()
}