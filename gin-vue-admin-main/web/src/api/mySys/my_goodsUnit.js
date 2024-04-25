import service from '@/utils/request'

// @Tags My_goodsUnit
// @Summary 创建计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goodsUnit true "创建计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /my_goodsUnit/createMy_goodsUnit [post]
export const createMy_goodsUnit = (data) => {
  return service({
    url: '/my_goodsUnit/createMy_goodsUnit',
    method: 'post',
    data
  })
}

// @Tags My_goodsUnit
// @Summary 删除计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goodsUnit true "删除计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goodsUnit/deleteMy_goodsUnit [delete]
export const deleteMy_goodsUnit = (params) => {
  return service({
    url: '/my_goodsUnit/deleteMy_goodsUnit',
    method: 'delete',
    params
  })
}

// @Tags My_goodsUnit
// @Summary 批量删除计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /my_goodsUnit/deleteMy_goodsUnit [delete]
export const deleteMy_goodsUnitByIds = (params) => {
  return service({
    url: '/my_goodsUnit/deleteMy_goodsUnitByIds',
    method: 'delete',
    params
  })
}

// @Tags My_goodsUnit
// @Summary 更新计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.My_goodsUnit true "更新计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /my_goodsUnit/updateMy_goodsUnit [put]
export const updateMy_goodsUnit = (data) => {
  return service({
    url: '/my_goodsUnit/updateMy_goodsUnit',
    method: 'put',
    data
  })
}

// @Tags My_goodsUnit
// @Summary 用id查询计量单位
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.My_goodsUnit true "用id查询计量单位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /my_goodsUnit/findMy_goodsUnit [get]
export const findMy_goodsUnit = (params) => {
  return service({
    url: '/my_goodsUnit/findMy_goodsUnit',
    method: 'get',
    params
  })
}

// @Tags My_goodsUnit
// @Summary 分页获取计量单位列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取计量单位列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /my_goodsUnit/getMy_goodsUnitList [get]
export const getMy_goodsUnitList = (params) => {
  return service({
    url: '/my_goodsUnit/getMy_goodsUnitList',
    method: 'get',
    params
  })
}
