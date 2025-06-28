package router

import (
	"boxdb/controllers/user"

	"github.com/gin-gonic/gin"
)

func initUserController(group *gin.RouterGroup) {
	userController := user.UserController{}

	userGroup := group.Group("/user")

	{
		userGroup.POST("/login/phone", userController.LoginWithVerificationCode)
		userGroup.POST("/login/username", userController.LoginWithUsernamePassword)
		userGroup.POST("/info/add", userController.AdminAddUser)
		userGroup.POST("/info/update", userController.AdminUpdateUserInfo)
		userGroup.POST("/captcha/generate", userController.GenerateCaptcha)
	}
}
