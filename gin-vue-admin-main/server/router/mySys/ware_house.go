package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WareHouseRouter struct {
}

// InitWareHouseRouter 初始化 仓库 路由信息
func (s *WareHouseRouter) InitWareHouseRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wareHouseRouter := Router.Group("wareHouse").Use(middleware.OperationRecord())
	wareHouseRouterWithoutRecord := Router.Group("wareHouse")
	wareHouseRouterWithoutAuth := PublicRouter.Group("wareHouse")

	var wareHouseApi = v1.ApiGroupApp.MySysApiGroup.WareHouseApi
	{
		wareHouseRouter.POST("createWareHouse", wareHouseApi.CreateWareHouse)   // 新建仓库
		wareHouseRouter.DELETE("deleteWareHouse", wareHouseApi.DeleteWareHouse) // 删除仓库
		wareHouseRouter.DELETE("deleteWareHouseByIds", wareHouseApi.DeleteWareHouseByIds) // 批量删除仓库
		wareHouseRouter.PUT("updateWareHouse", wareHouseApi.UpdateWareHouse)    // 更新仓库
	}
	{
		wareHouseRouterWithoutRecord.GET("findWareHouse", wareHouseApi.FindWareHouse)        // 根据ID获取仓库
		wareHouseRouterWithoutRecord.GET("getWareHouseList", wareHouseApi.GetWareHouseList)  // 获取仓库列表
	}
	{
	    wareHouseRouterWithoutAuth.GET("getWareHousePublic", wareHouseApi.GetWareHousePublic)  // 获取仓库列表
	}
}
