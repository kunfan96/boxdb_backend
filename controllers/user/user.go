package user

import (
	"fmt"
	"net/http"
	"time"

	"boxdb/config"
	"boxdb/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 手机验证码登录
func (u *UserController) LoginWithVerificationCode(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 账号+密码+图形验证码登录
func (u *UserController) LoginWithUsernamePassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 管理员添加用户
func (u *UserController) AdminAddUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 管理员修改用户信息
func (u *UserController) AdminUpdateUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 生成二维码
func (u *UserController) GenerateCaptcha(c *gin.Context) {
	id, image, answer := utils.GenerateCaptcha()

	// set captcha cache in redis for 120s
	utils.SetRedisStringByKey(fmt.Sprintf("%s:%s", config.CAPTCHA_PREFIX, id), answer, time.Second*120)

	c.JSON(200, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"id":    id,
			"image": image,
		},
	})
}
