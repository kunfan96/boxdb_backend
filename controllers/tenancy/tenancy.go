package tenancy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TenancyController struct{}

// 租户列表
func (t *TenancyController) TenancyList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 租户下的用户列表
func (t *TenancyController) TenancyListUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 管理员添加租户
func (t *TenancyController) AdminAddTenancy(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 管理员修改租户信息
func (t *TenancyController) AdminUpdateTenancyInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 关联租户和用户
func (t *TenancyController) AdminRelateTenancyUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}

// 删除用户租户关联信息
func (t *TenancyController) AdminDeleteTenancyUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
	})
}
