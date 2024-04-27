package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
    mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type WareHouseApi struct {
}

var wareHouseService = service.ServiceGroupApp.MySysServiceGroup.WareHouseService


// CreateWareHouse 创建仓库
// @Tags WareHouse
// @Summary 创建仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.WareHouse true "创建仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wareHouse/createWareHouse [post]
func (wareHouseApi *WareHouseApi) CreateWareHouse(c *gin.Context) {
	var wareHouse mySys.WareHouse
	err := c.ShouldBindJSON(&wareHouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wareHouse.CreatedBy = utils.GetUserID(c)

	if err := wareHouseService.CreateWareHouse(&wareHouse); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWareHouse 删除仓库
// @Tags WareHouse
// @Summary 删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.WareHouse true "删除仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wareHouse/deleteWareHouse [delete]
func (wareHouseApi *WareHouseApi) DeleteWareHouse(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := wareHouseService.DeleteWareHouse(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWareHouseByIds 批量删除仓库
// @Tags WareHouse
// @Summary 批量删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wareHouse/deleteWareHouseByIds [delete]
func (wareHouseApi *WareHouseApi) DeleteWareHouseByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := wareHouseService.DeleteWareHouseByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWareHouse 更新仓库
// @Tags WareHouse
// @Summary 更新仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.WareHouse true "更新仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wareHouse/updateWareHouse [put]
func (wareHouseApi *WareHouseApi) UpdateWareHouse(c *gin.Context) {
	var wareHouse mySys.WareHouse
	err := c.ShouldBindJSON(&wareHouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    wareHouse.UpdatedBy = utils.GetUserID(c)

	if err := wareHouseService.UpdateWareHouse(wareHouse); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWareHouse 用id查询仓库
// @Tags WareHouse
// @Summary 用id查询仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.WareHouse true "用id查询仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wareHouse/findWareHouse [get]
func (wareHouseApi *WareHouseApi) FindWareHouse(c *gin.Context) {
	ID := c.Query("ID")
	if rewareHouse, err := wareHouseService.GetWareHouse(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewareHouse": rewareHouse}, c)
	}
}

// GetWareHouseList 分页获取仓库列表
// @Tags WareHouse
// @Summary 分页获取仓库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.WareHouseSearch true "分页获取仓库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wareHouse/getWareHouseList [get]
func (wareHouseApi *WareHouseApi) GetWareHouseList(c *gin.Context) {
	var pageInfo mySysReq.WareHouseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wareHouseService.GetWareHouseInfoList(pageInfo); err != nil {
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

// GetWareHousePublic 不需要鉴权的仓库接口
// @Tags WareHouse
// @Summary 不需要鉴权的仓库接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.WareHouseSearch true "分页获取仓库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wareHouse/getWareHouseList [get]
func (wareHouseApi *WareHouseApi) GetWareHousePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的仓库接口信息",
    }, "获取成功", c)
}
