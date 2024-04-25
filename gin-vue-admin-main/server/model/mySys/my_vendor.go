// 自动生成模板My_vendor
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 供货商管理 结构体  My_vendor
type My_vendor struct {
 global.GVA_MODEL 
      VenderName  string `json:"venderName" form:"venderName" gorm:"column:vender_name;comment:厂家名;" binding:"required"`  //厂家名 
      VenderAddr  string `json:"venderAddr" form:"venderAddr" gorm:"column:vender_addr;comment:厂家地址;" binding:"required"`  //厂家地址 
      Header  string `json:"header" form:"header" gorm:"column:header;comment:负责人姓名;" binding:"required"`  //负责人 
      VenderPhone  string `json:"venderPhone" form:"venderPhone" gorm:"column:vender_phone;comment:厂家电话;" binding:"required"`  //厂家电话 
      VenderClass  string `json:"venderClass" form:"venderClass" gorm:"column:vender_class;comment:厂家类型;" binding:"required"`  //厂家类型 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 供货商管理 My_vendor自定义表名 my_vendor
func (My_vendor) TableName() string {
  return "my_vendor"
}

