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

type CustomerApi struct {
}

var customerService = service.ServiceGroupApp.MySysServiceGroup.CustomerService


// CreateCustomer 创建顾客
// @Tags Customer
// @Summary 创建顾客
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.Customer true "创建顾客"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customer/createCustomer [post]
func (customerApi *CustomerApi) CreateCustomer(c *gin.Context) {
	var customer mySys.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    customer.CreatedBy = utils.GetUserID(c)

	if err := customerService.CreateCustomer(&customer); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCustomer 删除顾客
// @Tags Customer
// @Summary 删除顾客
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.Customer true "删除顾客"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/deleteCustomer [delete]
func (customerApi *CustomerApi) DeleteCustomer(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := customerService.DeleteCustomer(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCustomerByIds 批量删除顾客
// @Tags Customer
// @Summary 批量删除顾客
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /customer/deleteCustomerByIds [delete]
func (customerApi *CustomerApi) DeleteCustomerByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := customerService.DeleteCustomerByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCustomer 更新顾客
// @Tags Customer
// @Summary 更新顾客
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mySys.Customer true "更新顾客"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/updateCustomer [put]
func (customerApi *CustomerApi) UpdateCustomer(c *gin.Context) {
	var customer mySys.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    customer.UpdatedBy = utils.GetUserID(c)

	if err := customerService.UpdateCustomer(customer); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCustomer 用id查询顾客
// @Tags Customer
// @Summary 用id查询顾客
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySys.Customer true "用id查询顾客"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customer/findCustomer [get]
func (customerApi *CustomerApi) FindCustomer(c *gin.Context) {
	ID := c.Query("ID")
	if recustomer, err := customerService.GetCustomer(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recustomer": recustomer}, c)
	}
}

// GetCustomerList 分页获取顾客列表
// @Tags Customer
// @Summary 分页获取顾客列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.CustomerSearch true "分页获取顾客列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getCustomerList [get]
func (customerApi *CustomerApi) GetCustomerList(c *gin.Context) {
	var pageInfo mySysReq.CustomerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := customerService.GetCustomerInfoList(pageInfo); err != nil {
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

// GetCustomerPublic 不需要鉴权的顾客接口
// @Tags Customer
// @Summary 不需要鉴权的顾客接口
// @accept application/json
// @Produce application/json
// @Param data query mySysReq.CustomerSearch true "分页获取顾客列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getCustomerList [get]
func (customerApi *CustomerApi) GetCustomerPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的顾客接口信息",
    }, "获取成功", c)
}
