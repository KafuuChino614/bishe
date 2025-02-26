import service from '@/utils/request'

// @Tags Order
// @Summary 创建订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "创建订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /order/createOrder [post]
export const createOrder = (data) => {
  return service({
    url: '/order/createOrder',
    method: 'post',
    data
  })
}

// @Tags Order
// @Summary 删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
export const deleteOrder = (params) => {
  return service({
    url: '/order/deleteOrder',
    method: 'delete',
    params
  })
}

// @Tags Order
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
export const deleteOrderByIds = (params) => {
  return service({
    url: '/order/deleteOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags Order
// @Summary 更新订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "更新订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
export const updateOrder = (data) => {
  return service({
    url: '/order/updateOrder',
    method: 'put',
    data
  })
}

// @Tags Order
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Order true "用id查询订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
export const findOrder = (params) => {
  return service({
    url: '/order/findOrder',
    method: 'get',
    params
  })
}

// @Tags Order
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
export const getOrderList = (params) => {
  return service({
    url: '/order/getOrderList',
    method: 'get',
    params
  })
}

export const getetOrderPublic = (params) => {
  return service({
    url: 'http://localhost:8080/api/order/getOrderPublic',
    method: 'get',
    params
  })
}

export const getAllOrderProfitPublic = (params) => {
  return service({
    url: 'http://localhost:8080/api/order/getAllOrderProfitPublic',
    method: 'get',
    params
  })
}
export const getOrderGPTPublic = (params) => {
  return service({
    url: 'http://localhost:8080/api/order/getOrderGPTPublic',
    method: 'get',
    params
  })
}

export const getGoodsSalesPublic = (params) => {
  return service({
    url: 'http://localhost:8080/api/order/getGoodsSalesPublic',
    method: 'get',
    params
  })
}

export const getGoodsProfitPublic = (params) => {
  return service({
    url: 'http://localhost:8080/api/order/getGoodsProfitPublic',
    method: 'get',
    params
  })
}