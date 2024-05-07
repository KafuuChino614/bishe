package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

// InitCustomerRouter 初始化 顾客 路由信息
func (s *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	customerRouterWithoutAuth := PublicRouter.Group("customer")

	var customerApi = v1.ApiGroupApp.MySysApiGroup.CustomerApi
	{
		customerRouter.POST("createCustomer", customerApi.CreateCustomer)   // 新建顾客
		customerRouter.DELETE("deleteCustomer", customerApi.DeleteCustomer) // 删除顾客
		customerRouter.DELETE("deleteCustomerByIds", customerApi.DeleteCustomerByIds) // 批量删除顾客
		customerRouter.PUT("updateCustomer", customerApi.UpdateCustomer)    // 更新顾客
	}
	{
		customerRouterWithoutRecord.GET("findCustomer", customerApi.FindCustomer)        // 根据ID获取顾客
		customerRouterWithoutRecord.GET("getCustomerList", customerApi.GetCustomerList)  // 获取顾客列表
	}
	{
	    customerRouterWithoutAuth.GET("getCustomerPublic", customerApi.GetCustomerPublic)  // 获取顾客列表
	}
}
