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
                      GoodsName  string `json:"goodsName" form:"goodsName" `
                      GoodsPrice  *float64 `json:"goodsPrice" form:"goodsPrice" `
                      Discount  *float64 `json:"discount" form:"discount" `
    request.PageInfo
}
