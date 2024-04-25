package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type My_goodsUnitRouter struct {
}

// InitMy_goodsUnitRouter 初始化 计量单位 路由信息
func (s *My_goodsUnitRouter) InitMy_goodsUnitRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	my_goodsUnitRouter := Router.Group("my_goodsUnit").Use(middleware.OperationRecord())
	my_goodsUnitRouterWithoutRecord := Router.Group("my_goodsUnit")
	my_goodsUnitRouterWithoutAuth := PublicRouter.Group("my_goodsUnit")

	var my_goodsUnitApi = v1.ApiGroupApp.MySysApiGroup.My_goodsUnitApi
	{
		my_goodsUnitRouter.POST("createMy_goodsUnit", my_goodsUnitApi.CreateMy_goodsUnit)   // 新建计量单位
		my_goodsUnitRouter.DELETE("deleteMy_goodsUnit", my_goodsUnitApi.DeleteMy_goodsUnit) // 删除计量单位
		my_goodsUnitRouter.DELETE("deleteMy_goodsUnitByIds", my_goodsUnitApi.DeleteMy_goodsUnitByIds) // 批量删除计量单位
		my_goodsUnitRouter.PUT("updateMy_goodsUnit", my_goodsUnitApi.UpdateMy_goodsUnit)    // 更新计量单位
	}
	{
		my_goodsUnitRouterWithoutRecord.GET("findMy_goodsUnit", my_goodsUnitApi.FindMy_goodsUnit)        // 根据ID获取计量单位
		my_goodsUnitRouterWithoutRecord.GET("getMy_goodsUnitList", my_goodsUnitApi.GetMy_goodsUnitList)  // 获取计量单位列表
	}
	{
	    my_goodsUnitRouterWithoutAuth.GET("getMy_goodsUnitPublic", my_goodsUnitApi.GetMy_goodsUnitPublic)  // 获取计量单位列表
	}
}
