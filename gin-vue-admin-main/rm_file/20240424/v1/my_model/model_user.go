package my_model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/my_model"
    my_modelReq "github.com/flipped-aurora/gin-vue-admin/server/model/my_model/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type Model_userApi struct {
}

var model_userService = service.ServiceGroupApp.My_modelServiceGroup.Model_userService


// CreateModel_user 创建用户信息
// @Tags Model_user
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_model.Model_user true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /model_user/createModel_user [post]
func (model_userApi *Model_userApi) CreateModel_user(c *gin.Context) {
	var model_user my_model.Model_user
	err := c.ShouldBindJSON(&model_user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    model_user.CreatedBy = utils.GetUserID(c)

	if err := model_userService.CreateModel_user(&model_user); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteModel_user 删除用户信息
// @Tags Model_user
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_model.Model_user true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /model_user/deleteModel_user [delete]
func (model_userApi *Model_userApi) DeleteModel_user(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := model_userService.DeleteModel_user(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteModel_userByIds 批量删除用户信息
// @Tags Model_user
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /model_user/deleteModel_userByIds [delete]
func (model_userApi *Model_userApi) DeleteModel_userByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := model_userService.DeleteModel_userByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateModel_user 更新用户信息
// @Tags Model_user
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_model.Model_user true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /model_user/updateModel_user [put]
func (model_userApi *Model_userApi) UpdateModel_user(c *gin.Context) {
	var model_user my_model.Model_user
	err := c.ShouldBindJSON(&model_user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    model_user.UpdatedBy = utils.GetUserID(c)

	if err := model_userService.UpdateModel_user(model_user); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindModel_user 用id查询用户信息
// @Tags Model_user
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query my_model.Model_user true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /model_user/findModel_user [get]
func (model_userApi *Model_userApi) FindModel_user(c *gin.Context) {
	ID := c.Query("ID")
	if remodel_user, err := model_userService.GetModel_user(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remodel_user": remodel_user}, c)
	}
}

// GetModel_userList 分页获取用户信息列表
// @Tags Model_user
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query my_modelReq.Model_userSearch true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /model_user/getModel_userList [get]
func (model_userApi *Model_userApi) GetModel_userList(c *gin.Context) {
	var pageInfo my_modelReq.Model_userSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := model_userService.GetModel_userInfoList(pageInfo); err != nil {
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

// GetModel_userPublic 不需要鉴权的用户信息接口
// @Tags Model_user
// @Summary 不需要鉴权的用户信息接口
// @accept application/json
// @Produce application/json
// @Param data query my_modelReq.Model_userSearch true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /model_user/getModel_userList [get]
func (model_userApi *Model_userApi) GetModel_userPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户信息接口信息",
    }, "获取成功", c)
}
