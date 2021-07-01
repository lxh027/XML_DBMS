package routes

import (
	"github.com/gin-gonic/gin"
	"lxh027.com/xml-dbms/api-web/controller"
	"net/http"
)

func Routes(router *gin.Engine)  {
	router.POST("/access", controller.Auth)
	router.POST("/logout", controller.Logout)

	router.POST("/exeSql", controller.ExecSql)

	router.StaticFS("/sql/", http.Dir("./api-web/web"))
}
