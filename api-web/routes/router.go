package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine)  {

	router.StaticFS("/admin/", http.Dir("./web"))
}
