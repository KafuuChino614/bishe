package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WareHouseInfoRouter struct {
}

// InitWareHouseInfoRouter 初始化 仓库信息 路由信息
func (s *WareHouseInfoRouter) InitWareHouseInfoRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wareHouseInfoRouter := Router.Group("wareHouseInfo").Use(middleware.OperationRecord())
	wareHouseInfoRouterWithoutRecord := Router.Group("wareHouseInfo")
	wareHouseInfoRouterWithoutAuth := PublicRouter.Group("wareHouseInfo")

	var wareHouseInfoApi = v1.ApiGroupApp.MySysApiGroup.WareHouseInfoApi
	{
		wareHouseInfoRouter.POST("createWareHouseInfo", wareHouseInfoApi.CreateWareHouseInfo)   // 新建仓库信息
		wareHouseInfoRouter.DELETE("deleteWareHouseInfo", wareHouseInfoApi.DeleteWareHouseInfo) // 删除仓库信息
		wareHouseInfoRouter.DELETE("deleteWareHouseInfoByIds", wareHouseInfoApi.DeleteWareHouseInfoByIds) // 批量删除仓库信息
		wareHouseInfoRouter.PUT("updateWareHouseInfo", wareHouseInfoApi.UpdateWareHouseInfo)    // 更新仓库信息
	}
	{
		wareHouseInfoRouterWithoutRecord.GET("findWareHouseInfo", wareHouseInfoApi.FindWareHouseInfo)        // 根据ID获取仓库信息
		wareHouseInfoRouterWithoutRecord.GET("getWareHouseInfoList", wareHouseInfoApi.GetWareHouseInfoList)  // 获取仓库信息列表
	}
	{
	    wareHouseInfoRouterWithoutAuth.GET("getWareHouseInfoPublic", wareHouseInfoApi.GetWareHouseInfoPublic)  // 获取仓库信息列表
	}
}
