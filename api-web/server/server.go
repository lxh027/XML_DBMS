package server

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"lxh027.com/xml-dbms/api-web/routes"
	"lxh027.com/xml-dbms/config"
)

func Run(httpServer *gin.Engine){

	// 设置运行模式
	gin.SetMode(config.ApiServerConfig.Gin.Mode)
	httpServer = gin.Default()

	// 创建session存储引擎
	sessionStore := cookie.NewStore([]byte(config.ApiServerConfig.Session.Key))
	sessionStore.Options(sessions.Options{
		MaxAge:	config.ApiServerConfig.Session.Age,
		Path: config.ApiServerConfig.Session.Path,
	})
	httpServer.Use(sessions.Sessions(config.ApiServerConfig.Session.Name, sessionStore))

	// 设置恢复模式
	httpServer.Use(gin.Recovery())

	routes.Routes(httpServer)

	serverAddr := fmt.Sprintf("%s:%s", config.ApiServerConfig.Gin.Host, config.ApiServerConfig.Gin.Port)
	if err := httpServer.Run(serverAddr); err != nil {
		log.Panicf("start httpserver error: %v\n", err.Error())
	}
}