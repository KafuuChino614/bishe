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

type My_goodsApi struct {
}

var my_goodsService = service.ServiceGroupApp.MySysServiceGroup.My_goodsService


// CreateMy_goods 创建商品信息
// @Tags My_goods
// @Summary 创建商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goods true "创建商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_goods/createMy_goods [post]
func (my_goodsApi *My_goodsApi) CreateMy_goods(c *gin.Context) {
	var my_goods mySys.My_goods
	err := c.ShouldBindJSON(&my_goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    my_goods.CreatedBy = utils.GetUserID(c)

	if err := my_goodsService.CreateMy_goods(&my_goods); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMy_goods 删除商品信息
// @Tags My_goods
// @Summary 删除商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goods true "删除商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goods/deleteMy_goods [delete]
func (my_goodsApi *My_goodsApi) DeleteMy_goods(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := my_goodsService.DeleteMy_goods(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMy_goodsByIds 批量删除商品信息
// @Tags My_goods
// @Summary 批量删除商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /my_goods/deleteMy_goodsByIds [delete]
func (my_goodsApi *My_goodsApi) DeleteMy_goodsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := my_goodsService.DeleteMy_goodsByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMy_goods 更新商品信息
// @Tags My_goods
// @Summary 更新商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.My_goods true "更新商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_goods/updateMy_goods [put]
func (my_goodsApi *My_goodsApi) UpdateMy_goods(c *gin.Context) {
	var my_goods mySys.My_goods
	err := c.ShouldBindJSON(&my_goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    my_goods.UpdatedBy = utils.GetUserID(c)

	if err := my_goodsService.UpdateMy_goods(my_goods); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMy_goods 用id查询商品信息
// @Tags My_goods
// @Summary 用id查询商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.My_goods true "用id查询商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_goods/findMy_goods [get]
func (my_goodsApi *My_goodsApi) FindMy_goods(c *gin.Context) {
	ID := c.Query("ID")
	if remy_goods, err := my_goodsService.GetMy_goods(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remy_goods": remy_goods}, c)
	}
}

// GetMy_goodsList 分页获取商品信息列表
// @Tags My_goods
// @Summary 分页获取商品信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_goodsSearch true "分页获取商品信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goods/getMy_goodsList [get]
func (my_goodsApi *My_goodsApi) GetMy_goodsList(c *gin.Context) {
	var pageInfo mySysReq.My_goodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := my_goodsService.GetMy_goodsInfoList(pageInfo); err != nil {
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

// GetMy_goodsPublic 不需要鉴权的商品信息接口
// @Tags My_goods
// @Summary 不需要鉴权的商品信息接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.My_goodsSearch true "分页获取商品信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goods/getMy_goodsList [get]
func (my_goodsApi *My_goodsApi) GetMy_goodsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的商品信息接口信息",
    }, "获取成功", c)
}
