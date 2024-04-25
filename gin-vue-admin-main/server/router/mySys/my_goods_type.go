package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type My_goodsTypeRouter struct {
}

// InitMy_goodsTypeRouter 初始化 商品类型 路由信息
func (s *My_goodsTypeRouter) InitMy_goodsTypeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	my_goodsTypeRouter := Router.Group("my_goodsType").Use(middleware.OperationRecord())
	my_goodsTypeRouterWithoutRecord := Router.Group("my_goodsType")
	my_goodsTypeRouterWithoutAuth := PublicRouter.Group("my_goodsType")

	var my_goodsTypeApi = v1.ApiGroupApp.MySysApiGroup.My_goodsTypeApi
	{
		my_goodsTypeRouter.POST("createMy_goodsType", my_goodsTypeApi.CreateMy_goodsType)   // 新建商品类型
		my_goodsTypeRouter.DELETE("deleteMy_goodsType", my_goodsTypeApi.DeleteMy_goodsType) // 删除商品类型
		my_goodsTypeRouter.DELETE("deleteMy_goodsTypeByIds", my_goodsTypeApi.DeleteMy_goodsTypeByIds) // 批量删除商品类型
		my_goodsTypeRouter.PUT("updateMy_goodsType", my_goodsTypeApi.UpdateMy_goodsType)    // 更新商品类型
	}
	{
		my_goodsTypeRouterWithoutRecord.GET("findMy_goodsType", my_goodsTypeApi.FindMy_goodsType)        // 根据ID获取商品类型
		my_goodsTypeRouterWithoutRecord.GET("getMy_goodsTypeList", my_goodsTypeApi.GetMy_goodsTypeList)  // 获取商品类型列表
	}
	{
	    my_goodsTypeRouterWithoutAuth.GET("getMy_goodsTypePublic", my_goodsTypeApi.GetMy_goodsTypePublic)  // 获取商品类型列表
	}
}
