package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type OrderSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Uuid  string `json:"uuid" form:"uuid" `
                      CustomerName  string `json:"customer" form:"customer" `
                      WareHouseName  string `json:"wareHouseName" form:"wareHouseName" `
                      GoodsName  string `json:"goodsName" form:"goodsName" `
                      GoodsType  string `json:"goodsType" form:"goodsType" `
                      GoodsPrice  *float64 `json:"goodsPrice" form:"goodsPrice" `
                      Discount  string `json:"discount" form:"discount" `
    request.PageInfo
}
