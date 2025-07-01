package middleware

import (
	"fmt"
	"net/http"
	"slices"

	"boxdb/config"
	"boxdb/utils"

	"github.com/gin-gonic/gin"
)

func CheckAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(config.SKIP_AUTH_API, c.FullPath()) {
			c.Next()
		} else {
			token := c.Request.Header.Get(config.BOXDB_TOKEN_KEY)

			str, _ := utils.GetRedisStringByKey(fmt.Sprintf("%s:%s", config.LOGIN_USER_TOKEN_PREFIX, token))
			if str != "" {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusUnauthorized,
					"msg":  "暂无权限",
					"data": nil,
				})

				c.Abort()
			}

		}

	}
}
