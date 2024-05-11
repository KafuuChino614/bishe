package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

// InitOrderRouter 初始化 订单 路由信息
func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	orderRouter := Router.Group("order").Use(middleware.OperationRecord())
	orderRouterWithoutRecord := Router.Group("order")
	orderRouterWithoutAuth := PublicRouter.Group("order")

	var orderApi = v1.ApiGroupApp.MySysApiGroup.OrderApi
	{
		orderRouter.POST("createOrder", orderApi.CreateOrder)             // 新建订单
		orderRouter.DELETE("deleteOrder", orderApi.DeleteOrder)           // 删除订单
		orderRouter.DELETE("deleteOrderByIds", orderApi.DeleteOrderByIds) // 批量删除订单
		orderRouter.PUT("updateOrder", orderApi.UpdateOrder)              // 更新订单
	}
	{
		orderRouterWithoutRecord.GET("findOrder", orderApi.FindOrder)       // 根据ID获取订单
		orderRouterWithoutRecord.GET("getOrderList", orderApi.GetOrderList) // 获取订单列表
	}
	{
		orderRouterWithoutAuth.GET("getOrderPublic", orderApi.GetOrderPublic)                   // 获取周销售
		orderRouterWithoutAuth.GET("getAllOrderProfitPublic", orderApi.GetAllOrderProfitPublic) // 获取利润数据
		orderRouterWithoutAuth.GET("getOrderGPTPublic", orderApi.GetOrderGPTPublic)             // 获取利润数据
	}
}
