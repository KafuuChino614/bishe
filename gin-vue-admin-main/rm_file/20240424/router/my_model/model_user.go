package my_model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Model_userRouter struct {
}

// InitModel_userRouter 初始化 用户信息 路由信息
func (s *Model_userRouter) InitModel_userRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	model_userRouter := Router.Group("model_user").Use(middleware.OperationRecord())
	model_userRouterWithoutRecord := Router.Group("model_user")
	model_userRouterWithoutAuth := PublicRouter.Group("model_user")

	var model_userApi = v1.ApiGroupApp.My_modelApiGroup.Model_userApi
	{
		model_userRouter.POST("createModel_user", model_userApi.CreateModel_user)   // 新建用户信息
		model_userRouter.DELETE("deleteModel_user", model_userApi.DeleteModel_user) // 删除用户信息
		model_userRouter.DELETE("deleteModel_userByIds", model_userApi.DeleteModel_userByIds) // 批量删除用户信息
		model_userRouter.PUT("updateModel_user", model_userApi.UpdateModel_user)    // 更新用户信息
	}
	{
		model_userRouterWithoutRecord.GET("findModel_user", model_userApi.FindModel_user)        // 根据ID获取用户信息
		model_userRouterWithoutRecord.GET("getModel_userList", model_userApi.GetModel_userList)  // 获取用户信息列表
	}
	{
	    model_userRouterWithoutAuth.GET("getModel_userPublic", model_userApi.GetModel_userPublic)  // 获取用户信息列表
	}
}
