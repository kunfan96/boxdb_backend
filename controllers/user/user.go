package user

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"boxdb/config"
	"boxdb/models"
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

// username+password+captcha login
func (u *UserController) LoginWithUsernamePassword(c *gin.Context) {
	reqBody := LoginWithUsernamePasswordReqBody{}
	c.ShouldBindBodyWithJSON(&reqBody)

	if reqBody.Captcha.Code == "" || reqBody.Captcha.Id == "" || reqBody.Username == "" || reqBody.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "请输入完整信息",
			"data": nil,
		})

		return
	}

	// base64 encode
	pwd, _ := base64.StdEncoding.DecodeString(reqBody.Password)
	fullPwd := string(pwd)
	// frontend use string like {6 chars}Boxdb_654321{4 chars}
	reqBody.Password = fullPwd[6 : len(fullPwd)-4]

	captcha, err := utils.GetRedisStringByKey(fmt.Sprintf("%s:%s", config.CAPTCHA_PREFIX, reqBody.Captcha.Id))

	if captcha != reqBody.Captcha.Code || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "验证码错误",
			"data": nil,
		})
		return
	}

	user := models.USER{}
	query := utils.DB.Where("name = ?", reqBody.Username).Find(&user)

	// get a record in database
	if query.Error == nil {
		// check two password
		if utils.CheckPassword(user.Password, reqBody.Password) {
			token, _ := utils.GenerateToken(18)
			c.Header("Token", token)

			// delete old user login token if exist
			oldToken, _ := utils.GetRedisStringByKey(fmt.Sprintf("%s:%s", config.LOGIN_USER_UID_PREFIX, user.ID))
			if oldToken != "" {
				utils.DelRedisStringByKey(fmt.Sprintf("%s:%s", config.LOGIN_USER_TOKEN_PREFIX, oldToken))
				utils.DelRedisStringByKey(fmt.Sprintf("%s:%s", config.LOGIN_USER_UID_PREFIX, user.ID))
			}

			// set token in redis
			utils.SetRedisStringByKey(fmt.Sprintf("%s:%s", config.LOGIN_USER_TOKEN_PREFIX, token), user.ID, time.Hour*72)
			utils.SetRedisStringByKey(fmt.Sprintf("%s:%s", config.LOGIN_USER_UID_PREFIX, user.ID), token, time.Hour*72)

			// password and username correct
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": nil,
			})
		} else {
			// password error
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "密码错误",
				"data": nil,
			})
		}
	} else {
		// username error
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "暂无该用户",
			"data": nil,
		})
	}
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

// 生成验证码
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
