// 自动生成模板My_goodsUnit
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 计量单位 结构体  My_goodsUnit
type My_goodsUnit struct {
 global.GVA_MODEL 
      UnitName  string `json:"unitName" form:"unitName" gorm:"column:unit_name;comment:计量单位;" binding:"required"`  //计量单位 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 计量单位 My_goodsUnit自定义表名 my_goodsUnit
func (My_goodsUnit) TableName() string {
  return "my_goodsUnit"
}

