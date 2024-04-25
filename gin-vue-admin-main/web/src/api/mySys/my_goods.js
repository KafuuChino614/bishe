import service from '@/utils/request'

// @Tags My_goods
// @Summary 创建商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goods true "创建商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_goods/createMy_goods [post]
export const createMy_goods = (data) => {
  return service({
    url: '/my_goods/createMy_goods',
    method: 'post',
    data
  })
}

// @Tags My_goods
// @Summary 删除商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goods true "删除商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goods/deleteMy_goods [delete]
export const deleteMy_goods = (params) => {
  return service({
    url: '/my_goods/deleteMy_goods',
    method: 'delete',
    params
  })
}

// @Tags My_goods
// @Summary 批量删除商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goods/deleteMy_goods [delete]
export const deleteMy_goodsByIds = (params) => {
  return service({
    url: '/my_goods/deleteMy_goodsByIds',
    method: 'delete',
    params
  })
}

// @Tags My_goods
// @Summary 更新商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goods true "更新商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_goods/updateMy_goods [put]
export const updateMy_goods = (data) => {
  return service({
    url: '/my_goods/updateMy_goods',
    method: 'put',
    data
  })
}

// @Tags My_goods
// @Summary 用id查询商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.My_goods true "用id查询商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_goods/findMy_goods [get]
export const findMy_goods = (params) => {
  return service({
    url: '/my_goods/findMy_goods',
    method: 'get',
    params
  })
}

// @Tags My_goods
// @Summary 分页获取商品信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goods/getMy_goodsList [get]
export const getMy_goodsList = (params) => {
  return service({
    url: '/my_goods/getMy_goodsList',
    method: 'get',
    params
  })
}
