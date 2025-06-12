package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册路由，由每个controller调用
func RegisterRouterGroupV1(group *gin.RouterGroup) {

	{
		initPing(group)

		initUserController(group)
		initTenancyController(group)
	}
}

// 测试心跳
func initPing(group *gin.RouterGroup) {
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": http.StatusOK,
			"msg":  "pong",
		})
	})
}
