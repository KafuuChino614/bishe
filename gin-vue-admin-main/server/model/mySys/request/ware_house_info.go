package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type WareHouseInfoSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      WareHouseName  string `json:"wareHouseName" form:"wareHouseName" `
                      GoodsName  string `json:"goodsName" form:"goodsName" `
                      GoodsType  string `json:"goodsType" form:"goodsType" `
    request.PageInfo
}
