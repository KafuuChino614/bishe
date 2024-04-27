// 自动生成模板WareHouse
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 仓库 结构体  WareHouse
type WareHouse struct {
 global.GVA_MODEL 
      WareHouseName  string `json:"wareHouseName" form:"wareHouseName" gorm:"column:ware_house_name;comment:仓库名字;" binding:"required"`  //仓库名字 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 仓库 WareHouse自定义表名 wareHouse
func (WareHouse) TableName() string {
  return "wareHouse"
}

