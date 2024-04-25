import service from '@/utils/request'

// @Tags My_user
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_user true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_user/createMy_user [post]
export const createMy_user = (data) => {
  return service({
    url: '/my_user/createMy_user',
    method: 'post',
    data
  })
}

// @Tags My_user
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_user true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_user/deleteMy_user [delete]
export const deleteMy_user = (params) => {
  return service({
    url: '/my_user/deleteMy_user',
    method: 'delete',
    params
  })
}

// @Tags My_user
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_user/deleteMy_user [delete]
export const deleteMy_userByIds = (params) => {
  return service({
    url: '/my_user/deleteMy_userByIds',
    method: 'delete',
    params
  })
}

// @Tags My_user
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_user true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_user/updateMy_user [put]
export const updateMy_user = (data) => {
  return service({
    url: '/my_user/updateMy_user',
    method: 'put',
    data
  })
}

// @Tags My_user
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.My_user true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_user/findMy_user [get]
export const findMy_user = (params) => {
  return service({
    url: '/my_user/findMy_user',
    method: 'get',
    params
  })
}

// @Tags My_user
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_user/getMy_userList [get]
export const getMy_userList = (params) => {
  return service({
    url: '/my_user/getMy_userList',
    method: 'get',
    params
  })
}
