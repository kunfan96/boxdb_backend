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
			fmt.Println(err.(string), 9999)
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

	BASE_ROUTER := gin.Default()
	API_V1_GROUP := BASE_ROUTER.Group("/boxdb/api/v1")
	router.RegisterRouterGroupV1(API_V1_GROUP)

	BASE_ROUTER.Run(fmt.Sprintf(":%s", port))
}
