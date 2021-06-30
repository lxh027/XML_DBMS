package controller

import (
	"github.com/gin-gonic/gin"
	"lxh027.com/xml-dbms/api-web/grpc_client"
	"lxh027.com/xml-dbms/api-web/helper"
	"lxh027.com/xml-dbms/constants"
	"net/http"
)

func ExecSql(c *gin.Context)  {
	// 检查登录
	if checkAccess(c) == false {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "未登录", nil))
		return
	}
	var sqlJson struct{
		Sql string `json:"sql"`
	}
	// 获取参数
	if err := c.ShouldBind(&sqlJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "参数错误", nil))
		return
	}
	// 访问db_server验证密码
	sqlResult, err := grpc_client.ExecSql(c, sqlJson.Sql)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "data server访问失败", err.Error()))
		return
	} else if sqlResult.Status != constants.SqlOK {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "执行失败", sqlResult.Message))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "执行成功", sqlResult))
}
