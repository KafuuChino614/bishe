import service from '@/utils/request'

// @Tags Model_user
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Model_user true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /model_user/createModel_user [post]
export const createModel_user = (data) => {
  return service({
    url: '/model_user/createModel_user',
    method: 'post',
    data
  })
}

// @Tags Model_user
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Model_user true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /model_user/deleteModel_user [delete]
export const deleteModel_user = (params) => {
  return service({
    url: '/model_user/deleteModel_user',
    method: 'delete',
    params
  })
}

// @Tags Model_user
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /model_user/deleteModel_user [delete]
export const deleteModel_userByIds = (params) => {
  return service({
    url: '/model_user/deleteModel_userByIds',
    method: 'delete',
    params
  })
}

// @Tags Model_user
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Model_user true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /model_user/updateModel_user [put]
export const updateModel_user = (data) => {
  return service({
    url: '/model_user/updateModel_user',
    method: 'put',
    data
  })
}

// @Tags Model_user
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Model_user true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /model_user/findModel_user [get]
export const findModel_user = (params) => {
  return service({
    url: '/model_user/findModel_user',
    method: 'get',
    params
  })
}

// @Tags Model_user
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /model_user/getModel_userList [get]
export const getModel_userList = (params) => {
  return service({
    url: '/model_user/getModel_userList',
    method: 'get',
    params
  })
}
