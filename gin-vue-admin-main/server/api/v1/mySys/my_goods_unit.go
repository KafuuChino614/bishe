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

type My_goodsUnitApi struct {
}

var my_goodsUnitService = service.ServiceGroupApp.MySysServiceGroup.My_goodsUnitService


// CreateMy_goodsUnit 创建计量单位
// @Tags My_goodsUnit
// @Summary 创建计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goodsUnit true "创建计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_goodsUnit/createMy_goodsUnit [post]
func (my_goodsUnitApi *My_goodsUnitApi) CreateMy_goodsUnit(c *gin.Context) {
	var my_goodsUnit mySys.My_goodsUnit
	err := c.ShouldBindJSON(&my_goodsUnit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    my_goodsUnit.CreatedBy = utils.GetUserID(c)

	if err := my_goodsUnitService.CreateMy_goodsUnit(&my_goodsUnit); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMy_goodsUnit 删除计量单位
// @Tags My_goodsUnit
// @Summary 删除计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goodsUnit true "删除计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goodsUnit/deleteMy_goodsUnit [delete]
func (my_goodsUnitApi *My_goodsUnitApi) DeleteMy_goodsUnit(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := my_goodsUnitService.DeleteMy_goodsUnit(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMy_goodsUnitByIds 批量删除计量单位
// @Tags My_goodsUnit
// @Summary 批量删除计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /my_goodsUnit/deleteMy_goodsUnitByIds [delete]
func (my_goodsUnitApi *My_goodsUnitApi) DeleteMy_goodsUnitByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := my_goodsUnitService.DeleteMy_goodsUnitByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMy_goodsUnit 更新计量单位
// @Tags My_goodsUnit
// @Summary 更新计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goodsUnit true "更新计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_goodsUnit/updateMy_goodsUnit [put]
func (my_goodsUnitApi *My_goodsUnitApi) UpdateMy_goodsUnit(c *gin.Context) {
	var my_goodsUnit mySys.My_goodsUnit
	err := c.ShouldBindJSON(&my_goodsUnit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    my_goodsUnit.UpdatedBy = utils.GetUserID(c)

	if err := my_goodsUnitService.UpdateMy_goodsUnit(my_goodsUnit); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMy_goodsUnit 用id查询计量单位
// @Tags My_goodsUnit
// @Summary 用id查询计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.My_goodsUnit true "用id查询计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_goodsUnit/findMy_goodsUnit [get]
func (my_goodsUnitApi *My_goodsUnitApi) FindMy_goodsUnit(c *gin.Context) {
	ID := c.Query("ID")
	if remy_goodsUnit, err := my_goodsUnitService.GetMy_goodsUnit(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remy_goodsUnit": remy_goodsUnit}, c)
	}
}

// GetMy_goodsUnitList 分页获取计量单位列表
// @Tags My_goodsUnit
// @Summary 分页获取计量单位列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_goodsUnitSearch true "分页获取计量单位列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goodsUnit/getMy_goodsUnitList [get]
func (my_goodsUnitApi *My_goodsUnitApi) GetMy_goodsUnitList(c *gin.Context) {
	var pageInfo mySysReq.My_goodsUnitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := my_goodsUnitService.GetMy_goodsUnitInfoList(pageInfo); err != nil {
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

// GetMy_goodsUnitPublic 不需要鉴权的计量单位接口
// @Tags My_goodsUnit
// @Summary 不需要鉴权的计量单位接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_goodsUnitSearch true "分页获取计量单位列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goodsUnit/getMy_goodsUnitList [get]
func (my_goodsUnitApi *My_goodsUnitApi) GetMy_goodsUnitPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的计量单位接口信息",
    }, "获取成功", c)
}
