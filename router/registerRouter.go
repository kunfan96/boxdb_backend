package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("1111")
}

// 注册路由，由每个controller调用
func RegisterRouter(group *gin.RouterGroup) {
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": http.StatusOK,
			"msg":  "pong",
		})
	})
}
