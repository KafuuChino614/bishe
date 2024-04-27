import service from '@/utils/request'

// @Tags WareHouse
// @Summary 创建仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WareHouse true "创建仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wareHouse/createWareHouse [post]
export const createWareHouse = (data) => {
  return service({
    url: '/wareHouse/createWareHouse',
    method: 'post',
    data
  })
}

// @Tags WareHouse
// @Summary 删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WareHouse true "删除仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wareHouse/deleteWareHouse [delete]
export const deleteWareHouse = (params) => {
  return service({
    url: '/wareHouse/deleteWareHouse',
    method: 'delete',
    params
  })
}

// @Tags WareHouse
// @Summary 批量删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wareHouse/deleteWareHouse [delete]
export const deleteWareHouseByIds = (params) => {
  return service({
    url: '/wareHouse/deleteWareHouseByIds',
    method: 'delete',
    params
  })
}

// @Tags WareHouse
// @Summary 更新仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WareHouse true "更新仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wareHouse/updateWareHouse [put]
export const updateWareHouse = (data) => {
  return service({
    url: '/wareHouse/updateWareHouse',
    method: 'put',
    data
  })
}

// @Tags WareHouse
// @Summary 用id查询仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WareHouse true "用id查询仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wareHouse/findWareHouse [get]
export const findWareHouse = (params) => {
  return service({
    url: '/wareHouse/findWareHouse',
    method: 'get',
    params
  })
}

// @Tags WareHouse
// @Summary 分页获取仓库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取仓库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wareHouse/getWareHouseList [get]
export const getWareHouseList = (params) => {
  return service({
    url: '/wareHouse/getWareHouseList',
    method: 'get',
    params
  })
}
