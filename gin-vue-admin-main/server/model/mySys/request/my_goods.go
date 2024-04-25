package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type My_goodsSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      GoodsName  string `json:"goodsName" form:"goodsName" `
                      GoodsType  string `json:"goodsType" form:"goodsType" `
    request.PageInfo
}
