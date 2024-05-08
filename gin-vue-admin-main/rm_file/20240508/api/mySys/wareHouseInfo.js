import service from '@/utils/request'

// @Tags WareHouseInfo
// @Summary 创建仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WareHouseInfo true "创建仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wareHouseInfo/createWareHouseInfo [post]
export const createWareHouseInfo = (data) => {
  return service({
    url: '/wareHouseInfo/createWareHouseInfo',
    method: 'post',
    data
  })
}

// @Tags WareHouseInfo
// @Summary 删除仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WareHouseInfo true "删除仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wareHouseInfo/deleteWareHouseInfo [delete]
export const deleteWareHouseInfo = (params) => {
  return service({
    url: '/wareHouseInfo/deleteWareHouseInfo',
    method: 'delete',
    params
  })
}

// @Tags WareHouseInfo
// @Summary 批量删除仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wareHouseInfo/deleteWareHouseInfo [delete]
export const deleteWareHouseInfoByIds = (params) => {
  return service({
    url: '/wareHouseInfo/deleteWareHouseInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags WareHouseInfo
// @Summary 更新仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WareHouseInfo true "更新仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wareHouseInfo/updateWareHouseInfo [put]
export const updateWareHouseInfo = (data) => {
  return service({
    url: '/wareHouseInfo/updateWareHouseInfo',
    method: 'put',
    data
  })
}

// @Tags WareHouseInfo
// @Summary 用id查询仓库信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WareHouseInfo true "用id查询仓库信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wareHouseInfo/findWareHouseInfo [get]
export const findWareHouseInfo = (params) => {
  return service({
    url: '/wareHouseInfo/findWareHouseInfo',
    method: 'get',
    params
  })
}

// @Tags WareHouseInfo
// @Summary 分页获取仓库信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取仓库信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wareHouseInfo/getWareHouseInfoList [get]
export const getWareHouseInfoList = (params) => {
  return service({
    url: '/wareHouseInfo/getWareHouseInfoList',
    method: 'get',
    params
  })
}
