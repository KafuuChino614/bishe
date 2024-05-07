// 自动生成模板Customer
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 顾客 结构体  Customer
type Customer struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:客户名;" binding:"required"`  //客户名 
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:电话;"`  //电话 
      Email  string `json:"email" form:"email" gorm:"column:email;comment:客户邮箱;"`  //邮箱 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 顾客 Customer自定义表名 customer
func (Customer) TableName() string {
  return "customer"
}

