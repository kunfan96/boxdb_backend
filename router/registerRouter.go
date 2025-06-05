package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var GroupV1 *gin.RouterGroup

// 注册路由，由每个controller调用
func RegisterRouterGroupV1(group *gin.RouterGroup) {
	GroupV1 = group

	initPing()
}

// 测试心跳
func initPing() {
	GroupV1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": http.StatusOK,
			"msg":  "pong",
		})
	})
}
