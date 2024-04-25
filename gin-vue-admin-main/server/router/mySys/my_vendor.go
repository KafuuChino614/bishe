package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type My_vendorRouter struct {
}

// InitMy_vendorRouter 初始化 供货商管理 路由信息
func (s *My_vendorRouter) InitMy_vendorRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	my_vendorRouter := Router.Group("my_vendor").Use(middleware.OperationRecord())
	my_vendorRouterWithoutRecord := Router.Group("my_vendor")
	my_vendorRouterWithoutAuth := PublicRouter.Group("my_vendor")

	var my_vendorApi = v1.ApiGroupApp.MySysApiGroup.My_vendorApi
	{
		my_vendorRouter.POST("createMy_vendor", my_vendorApi.CreateMy_vendor)   // 新建供货商管理
		my_vendorRouter.DELETE("deleteMy_vendor", my_vendorApi.DeleteMy_vendor) // 删除供货商管理
		my_vendorRouter.DELETE("deleteMy_vendorByIds", my_vendorApi.DeleteMy_vendorByIds) // 批量删除供货商管理
		my_vendorRouter.PUT("updateMy_vendor", my_vendorApi.UpdateMy_vendor)    // 更新供货商管理
	}
	{
		my_vendorRouterWithoutRecord.GET("findMy_vendor", my_vendorApi.FindMy_vendor)        // 根据ID获取供货商管理
		my_vendorRouterWithoutRecord.GET("getMy_vendorList", my_vendorApi.GetMy_vendorList)  // 获取供货商管理列表
	}
	{
	    my_vendorRouterWithoutAuth.GET("getMy_vendorPublic", my_vendorApi.GetMy_vendorPublic)  // 获取供货商管理列表
	}
}
