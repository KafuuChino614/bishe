import service from '@/utils/request'

// @Tags My_vendor
// @Summary 创建供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_vendor true "创建供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_vendor/createMy_vendor [post]
export const createMy_vendor = (data) => {
  return service({
    url: '/my_vendor/createMy_vendor',
    method: 'post',
    data
  })
}

// @Tags My_vendor
// @Summary 删除供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_vendor true "删除供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_vendor/deleteMy_vendor [delete]
export const deleteMy_vendor = (params) => {
  return service({
    url: '/my_vendor/deleteMy_vendor',
    method: 'delete',
    params
  })
}

// @Tags My_vendor
// @Summary 批量删除供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_vendor/deleteMy_vendor [delete]
export const deleteMy_vendorByIds = (params) => {
  return service({
    url: '/my_vendor/deleteMy_vendorByIds',
    method: 'delete',
    params
  })
}

// @Tags My_vendor
// @Summary 更新供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_vendor true "更新供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_vendor/updateMy_vendor [put]
export const updateMy_vendor = (data) => {
  return service({
    url: '/my_vendor/updateMy_vendor',
    method: 'put',
    data
  })
}

// @Tags My_vendor
// @Summary 用id查询供货商管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.My_vendor true "用id查询供货商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_vendor/findMy_vendor [get]
export const findMy_vendor = (params) => {
  return service({
    url: '/my_vendor/findMy_vendor',
    method: 'get',
    params
  })
}

// @Tags My_vendor
// @Summary 分页获取供货商管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取供货商管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_vendor/getMy_vendorList [get]
export const getMy_vendorList = (params) => {
  return service({
    url: '/my_vendor/getMy_vendorList',
    method: 'get',
    params
  })
}
