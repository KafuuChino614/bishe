package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type My_goodsRouter struct {
}

// InitMy_goodsRouter 初始化 商品信息 路由信息
func (s *My_goodsRouter) InitMy_goodsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	my_goodsRouter := Router.Group("my_goods").Use(middleware.OperationRecord())
	my_goodsRouterWithoutRecord := Router.Group("my_goods")
	my_goodsRouterWithoutAuth := PublicRouter.Group("my_goods")

	var my_goodsApi = v1.ApiGroupApp.MySysApiGroup.My_goodsApi
	{
		my_goodsRouter.POST("createMy_goods", my_goodsApi.CreateMy_goods)   // 新建商品信息
		my_goodsRouter.DELETE("deleteMy_goods", my_goodsApi.DeleteMy_goods) // 删除商品信息
		my_goodsRouter.DELETE("deleteMy_goodsByIds", my_goodsApi.DeleteMy_goodsByIds) // 批量删除商品信息
		my_goodsRouter.PUT("updateMy_goods", my_goodsApi.UpdateMy_goods)    // 更新商品信息
	}
	{
		my_goodsRouterWithoutRecord.GET("findMy_goods", my_goodsApi.FindMy_goods)        // 根据ID获取商品信息
		my_goodsRouterWithoutRecord.GET("getMy_goodsList", my_goodsApi.GetMy_goodsList)  // 获取商品信息列表
	}
	{
	    my_goodsRouterWithoutAuth.GET("getMy_goodsPublic", my_goodsApi.GetMy_goodsPublic)  // 获取商品信息列表
	}
}
