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

type My_vendorApi struct {
}

var my_vendorService = service.ServiceGroupApp.MySysServiceGroup.My_vendorService


// CreateMy_vendor 创建供货商管理
// @Tags My_vendor
// @Summary 创建供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_vendor true "创建供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_vendor/createMy_vendor [post]
func (my_vendorApi *My_vendorApi) CreateMy_vendor(c *gin.Context) {
	var my_vendor mySys.My_vendor
	err := c.ShouldBindJSON(&my_vendor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    my_vendor.CreatedBy = utils.GetUserID(c)

	if err := my_vendorService.CreateMy_vendor(&my_vendor); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMy_vendor 删除供货商管理
// @Tags My_vendor
// @Summary 删除供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_vendor true "删除供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_vendor/deleteMy_vendor [delete]
func (my_vendorApi *My_vendorApi) DeleteMy_vendor(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := my_vendorService.DeleteMy_vendor(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMy_vendorByIds 批量删除供货商管理
// @Tags My_vendor
// @Summary 批量删除供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /my_vendor/deleteMy_vendorByIds [delete]
func (my_vendorApi *My_vendorApi) DeleteMy_vendorByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := my_vendorService.DeleteMy_vendorByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMy_vendor 更新供货商管理
// @Tags My_vendor
// @Summary 更新供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_vendor true "更新供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_vendor/updateMy_vendor [put]
func (my_vendorApi *My_vendorApi) UpdateMy_vendor(c *gin.Context) {
	var my_vendor mySys.My_vendor
	err := c.ShouldBindJSON(&my_vendor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    my_vendor.UpdatedBy = utils.GetUserID(c)

	if err := my_vendorService.UpdateMy_vendor(my_vendor); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMy_vendor 用id查询供货商管理
// @Tags My_vendor
// @Summary 用id查询供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.My_vendor true "用id查询供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_vendor/findMy_vendor [get]
func (my_vendorApi *My_vendorApi) FindMy_vendor(c *gin.Context) {
	ID := c.Query("ID")
	if remy_vendor, err := my_vendorService.GetMy_vendor(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remy_vendor": remy_vendor}, c)
	}
}

// GetMy_vendorList 分页获取供货商管理列表
// @Tags My_vendor
// @Summary 分页获取供货商管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_vendorSearch true "分页获取供货商管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_vendor/getMy_vendorList [get]
func (my_vendorApi *My_vendorApi) GetMy_vendorList(c *gin.Context) {
	var pageInfo mySysReq.My_vendorSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := my_vendorService.GetMy_vendorInfoList(pageInfo); err != nil {
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

// GetMy_vendorPublic 不需要鉴权的供货商管理接口
// @Tags My_vendor
// @Summary 不需要鉴权的供货商管理接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_vendorSearch true "分页获取供货商管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_vendor/getMy_vendorList [get]
func (my_vendorApi *My_vendorApi) GetMy_vendorPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的供货商管理接口信息",
    }, "获取成功", c)
}
