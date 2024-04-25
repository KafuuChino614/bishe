// 自动生成模板My_goodsType
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 商品类型 结构体  My_goodsType
type My_goodsType struct {
 global.GVA_MODEL 
      TypeName  string `json:"typeName" form:"typeName" gorm:"column:type_name;comment:商品名;" binding:"required"`  //商品名 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 商品类型 My_goodsType自定义表名 my_goodsType
func (My_goodsType) TableName() string {
  return "my_goodsType"
}

