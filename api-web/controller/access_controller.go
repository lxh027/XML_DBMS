package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"lxh027.com/xml-dbms/api-web/grpc_client"
	"lxh027.com/xml-dbms/api-web/helper"
	"lxh027.com/xml-dbms/constants"
	"lxh027.com/xml-dbms/proto"
	"net/http"
)

func Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "注销成功", nil))
}

func Auth(c *gin.Context)  {
	// 检查登录
	if checkAccess(c) {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "已登录", nil))
		return
	}
	var accessJson struct{
		Password string `json:"password"`
	}
	// 获取参数
	if err := c.ShouldBind(accessJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "参数错误", nil))
		return
	}
	// 访问db_server验证密码
	authResponse, err := grpc_client.ExecAuth(c, accessJson.Password)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "data server访问失败", err.Error()))
		return
	} else if authResponse.Status != proto.AuthResponse_OK {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "验证失败", authResponse.Message))
		return
	}
	session := sessions.Default(c)
	session.Set("access", "ok")
	_ = session.Save()
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "验证成功", authResponse.Message))
}

//检查登录
func checkAccess(c *gin.Context) bool {
	session := sessions.Default(c)
	if access := session.Get("access"); access == nil {
		return false
	}
	return true
}

