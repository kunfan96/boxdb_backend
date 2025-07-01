package main

import (
	"boxdb/config"
	"boxdb/router"
	"boxdb/utils"
	"boxdb/utils/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			utils.Logger.Error(err.(string))
		}
	}()

	// check whether the port is used
	// if port is used, quit
	port := utils.GetBootstrapConfig().Service.Port
	if utils.IsPortUsed(port) {
		str := fmt.Sprintf("the port %s has been used", port)
		fmt.Println(str)
		panic(str)
	}

	{
		utils.InitRedis()
		utils.InitMysql()
	}

	baseRouter := gin.Default()
	groupV1 := baseRouter.Group(config.API_V1_PREFIX)

	groupV1.Use(middleware.CheckAuthToken())
	router.RegisterRouterGroupV1(groupV1)

	baseRouter.Run(fmt.Sprintf(":%s", port))
}
