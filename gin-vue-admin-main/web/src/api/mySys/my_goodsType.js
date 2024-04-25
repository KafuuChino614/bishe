import service from '@/utils/request'

// @Tags My_goodsType
// @Summary 创建商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goodsType true "创建商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_goodsType/createMy_goodsType [post]
export const createMy_goodsType = (data) => {
  return service({
    url: '/my_goodsType/createMy_goodsType',
    method: 'post',
    data
  })
}

// @Tags My_goodsType
// @Summary 删除商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goodsType true "删除商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goodsType/deleteMy_goodsType [delete]
export const deleteMy_goodsType = (params) => {
  return service({
    url: '/my_goodsType/deleteMy_goodsType',
    method: 'delete',
    params
  })
}

// @Tags My_goodsType
// @Summary 批量删除商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goodsType/deleteMy_goodsType [delete]
export const deleteMy_goodsTypeByIds = (params) => {
  return service({
    url: '/my_goodsType/deleteMy_goodsTypeByIds',
    method: 'delete',
    params
  })
}

// @Tags My_goodsType
// @Summary 更新商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goodsType true "更新商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_goodsType/updateMy_goodsType [put]
export const updateMy_goodsType = (data) => {
  return service({
    url: '/my_goodsType/updateMy_goodsType',
    method: 'put',
    data
  })
}

// @Tags My_goodsType
// @Summary 用id查询商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.My_goodsType true "用id查询商品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_goodsType/findMy_goodsType [get]
export const findMy_goodsType = (params) => {
  return service({
    url: '/my_goodsType/findMy_goodsType',
    method: 'get',
    params
  })
}

// @Tags My_goodsType
// @Summary 分页获取商品类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goodsType/getMy_goodsTypeList [get]
export const getMy_goodsTypeList = (params) => {
  return service({
    url: '/my_goodsType/getMy_goodsTypeList',
    method: 'get',
    params
  })
}
