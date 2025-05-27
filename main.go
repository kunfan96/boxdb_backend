package main

import (
	"boxdb/router"
	"boxdb/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			utils.Logger.Error(err.(string))
			utils.Logger.Info("898989")
		}
	}()

	port := utils.GetBootstrapConfig().Service.Port
	if utils.IsPortUsed(port) {
		panic(fmt.Sprintf("the port %s has been used", port))
	}

	BASE_ROUTER := gin.Default()
	API_V1 := BASE_ROUTER.Group("/boxdb/api/v1")
	router.RegisterRouter(API_V1)

	BASE_ROUTER.Run(fmt.Sprintf(":%s", port))
}
