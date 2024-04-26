package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
	mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type My_goodsTypeApi struct {
}

var my_goodsTypeService = service.ServiceGroupApp.MySysServiceGroup.My_goodsTypeService

// CreateMy_goodsType 创建商品类型
// @Tags My_goodsType
// @Summary 创建商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goodsType true "创建商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_goodsType/createMy_goodsType [post]
func (my_goodsTypeApi *My_goodsTypeApi) CreateMy_goodsType(c *gin.Context) {
	var my_goodsType mySys.My_goodsType
	err := c.ShouldBindJSON(&my_goodsType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	my_goodsType.CreatedBy = utils.GetUserID(c)

	if err := my_goodsTypeService.CreateMy_goodsType(&my_goodsType); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMy_goodsType 删除商品类型
// @Tags My_goodsType
// @Summary 删除商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goodsType true "删除商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goodsType/deleteMy_goodsType [delete]
func (my_goodsTypeApi *My_goodsTypeApi) DeleteMy_goodsType(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := my_goodsTypeService.DeleteMy_goodsType(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMy_goodsTypeByIds 批量删除商品类型
// @Tags My_goodsType
// @Summary 批量删除商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /my_goodsType/deleteMy_goodsTypeByIds [delete]
func (my_goodsTypeApi *My_goodsTypeApi) DeleteMy_goodsTypeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := my_goodsTypeService.DeleteMy_goodsTypeByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMy_goodsType 更新商品类型
// @Tags My_goodsType
// @Summary 更新商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goodsType true "更新商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_goodsType/updateMy_goodsType [put]
func (my_goodsTypeApi *My_goodsTypeApi) UpdateMy_goodsType(c *gin.Context) {
	var my_goodsType mySys.My_goodsType
	err := c.ShouldBindJSON(&my_goodsType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	my_goodsType.UpdatedBy = utils.GetUserID(c)

	if err := my_goodsTypeService.UpdateMy_goodsType(my_goodsType); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMy_goodsType 用id查询商品类型
// @Tags My_goodsType
// @Summary 用id查询商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.My_goodsType true "用id查询商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_goodsType/findMy_goodsType [get]
func (my_goodsTypeApi *My_goodsTypeApi) FindMy_goodsType(c *gin.Context) {
	ID := c.Query("ID")
	if remy_goodsType, err := my_goodsTypeService.GetMy_goodsType(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remy_goodsType": remy_goodsType}, c)
	}
}

// GetMy_goodsTypeList 分页获取商品类型列表
// @Tags My_goodsType
// @Summary 分页获取商品类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_goodsTypeSearch true "分页获取商品类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goodsType/getMy_goodsTypePublic [get]
func (my_goodsTypeApi *My_goodsTypeApi) GetMy_goodsTypeList(c *gin.Context) {
	var pageInfo mySysReq.My_goodsTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := my_goodsTypeService.GetMy_goodsTypeInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetMy_goodsTypePublic 不需要鉴权的商品类型接口
// @Tags My_goodsType
// @Summary 不需要鉴权的商品类型接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_goodsTypeSearch true "获取商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goodsType/GetMy_goodsTypePublic [get]
func (my_goodsTypeApi *My_goodsTypeApi) GetMy_goodsTypePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	// 获取所有商品类型
	var goodsTypeS []mySys.My_goodsType
	global.GVA_DB.Find(&goodsTypeS)
	response.OkWithDetailed(gin.H{
		"goodsTypes": goodsTypeS,
	}, "获取成功", c)
}
