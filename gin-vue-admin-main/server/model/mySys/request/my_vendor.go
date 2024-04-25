package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type My_vendorSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      VenderName  string `json:"venderName" form:"venderName" `
                      VenderAddr  string `json:"venderAddr" form:"venderAddr" `
                      Header  string `json:"header" form:"header" `
    request.PageInfo
}
