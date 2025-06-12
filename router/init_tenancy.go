package router

import (
	"boxdb/controllers/tenancy"

	"github.com/gin-gonic/gin"
)

func initTenancyController(group *gin.RouterGroup) {
	tenancyController := tenancy.TenancyController{}

	tenancyGroup := group.Group("/tenancy")

	{
		tenancyGroup.POST("/list", tenancyController.TenancyList)
		tenancyGroup.POST("/list/user", tenancyController.TenancyListUser)
		tenancyGroup.POST("/add", tenancyController.AdminAddTenancy)
		tenancyGroup.POST("/update", tenancyController.AdminUpdateTenancyInfo)
		tenancyGroup.POST("/relate-user", tenancyController.AdminRelateTenancyUser)
		tenancyGroup.POST("/delete-user", tenancyController.AdminDeleteTenancyUser)

	}
}
