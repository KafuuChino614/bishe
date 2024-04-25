package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type Model_userSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Name  string `json:"name" form:"name" `
                      Password  string `json:"password" form:"password" `
                      Auth  *int `json:"auth" form:"auth" `
    request.PageInfo
}
