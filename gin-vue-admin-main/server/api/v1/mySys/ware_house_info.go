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

type WareHouseInfoApi struct {
}

var wareHouseInfoService = service.ServiceGroupApp.MySysServiceGroup.WareHouseInfoService

// CreateWareHouseInfo 创建仓库信息
// @Tags WareHouseInfo
// @Summary 创建仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.WareHouseInfo true "创建仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wareHouseInfo/createWareHouseInfo [post]
func (wareHouseInfoApi *WareHouseInfoApi) CreateWareHouseInfo(c *gin.Context) {
	var wareHouseInfo mySys.WareHouseInfo
	err := c.ShouldBindJSON(&wareHouseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wareHouseInfo.CreatedBy = utils.GetUserID(c)

	if err := wareHouseInfoService.CreateWareHouseInfo(&wareHouseInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWareHouseInfo 删除仓库信息
// @Tags WareHouseInfo
// @Summary 删除仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.WareHouseInfo true "删除仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wareHouseInfo/deleteWareHouseInfo [delete]
func (wareHouseInfoApi *WareHouseInfoApi) DeleteWareHouseInfo(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := wareHouseInfoService.DeleteWareHouseInfo(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWareHouseInfoByIds 批量删除仓库信息
// @Tags WareHouseInfo
// @Summary 批量删除仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wareHouseInfo/deleteWareHouseInfoByIds [delete]
func (wareHouseInfoApi *WareHouseInfoApi) DeleteWareHouseInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := wareHouseInfoService.DeleteWareHouseInfoByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWareHouseInfo 更新仓库信息
// @Tags WareHouseInfo
// @Summary 更新仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.WareHouseInfo true "更新仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wareHouseInfo/updateWareHouseInfo [put]
func (wareHouseInfoApi *WareHouseInfoApi) UpdateWareHouseInfo(c *gin.Context) {
	var wareHouseInfo mySys.WareHouseInfo
	err := c.ShouldBindJSON(&wareHouseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wareHouseInfo.UpdatedBy = utils.GetUserID(c)

	if err := wareHouseInfoService.UpdateWareHouseInfo(wareHouseInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWareHouseInfo 用id查询仓库信息
// @Tags WareHouseInfo
// @Summary 用id查询仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.WareHouseInfo true "用id查询仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wareHouseInfo/findWareHouseInfo [get]
func (wareHouseInfoApi *WareHouseInfoApi) FindWareHouseInfo(c *gin.Context) {
	ID := c.Query("ID")
	if rewareHouseInfo, err := wareHouseInfoService.GetWareHouseInfo(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewareHouseInfo": rewareHouseInfo}, c)
	}
}

// GetWareHouseInfoList 分页获取仓库信息列表
// @Tags WareHouseInfo
// @Summary 分页获取仓库信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.WareHouseInfoSearch true "分页获取仓库信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wareHouseInfo/getWareHouseInfoList [get]
func (wareHouseInfoApi *WareHouseInfoApi) GetWareHouseInfoList(c *gin.Context) {
	var pageInfo mySysReq.WareHouseInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wareHouseInfoService.GetWareHouseInfoInfoList(pageInfo); err != nil {
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

// GetWareHouseInfoPublic 不需要鉴权的仓库信息接口
// @Tags WareHouseInfo
// @Summary 不需要鉴权的仓库信息接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.WareHouseInfoSearch true "分页获取仓库信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wareHouseInfo/getWareHouseInfoList [get]
func (wareHouseInfoApi *WareHouseInfoApi) GetWareHouseInfoPublic(c *gin.Context) {
	list := wareHouseInfoService.GetAllWareHouseInfo()
	// 创建嵌套映射用于存储统计结果
	wareHouseProductCounts := make(map[string]map[string]map[string]int)

	// 遍历列表并统计每个仓库中每种商品的数量
	for _, info := range list {
		// 获取仓库名称、商品类型和商品名称
		wareHouseName := info.WareHouseName
		goodsType := info.GoodsType
		goodsName := info.GoodsName
		// 检查外部映射中是否已存在仓库名称的键，如果不存在则创建内部映射
		if _, ok := wareHouseProductCounts[wareHouseName]; !ok {
			wareHouseProductCounts[wareHouseName] = make(map[string]map[string]int)
		}
		// 检查商品类型下是否已存在商品名称的键，如果不存在则创建内部映射
		if _, ok := wareHouseProductCounts[wareHouseName][goodsType]; !ok {
			wareHouseProductCounts[wareHouseName][goodsType] = make(map[string]int)
		}
		// 增加数量
		wareHouseProductCounts[wareHouseName][goodsType][goodsName] += *info.Num
	}
	response.OkWithDetailed(gin.H{
		"wareHouseProductCounts": wareHouseProductCounts,
	}, "获取成功", c)
}
