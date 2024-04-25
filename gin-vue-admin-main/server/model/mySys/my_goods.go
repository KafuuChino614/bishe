// 自动生成模板My_goods
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 商品信息 结构体  My_goods
type My_goods struct {
 global.GVA_MODEL 
      GoodsName  string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:;" binding:"required"`  //商品名 
      GoodsType  string `json:"goodsType" form:"goodsType" gorm:"column:goods_type;comment:;" binding:"required"`  //商品类型 
      GoodsUnit  string `json:"goodsUnit" form:"goodsUnit" gorm:"column:goods_unit;comment:;" binding:"required"`  //商品单位 
      GoodsPrice  *float64 `json:"goodsPrice" form:"goodsPrice" gorm:"column:goods_price;comment:;" binding:"required"`  //商品价格 
      GoodsNum  *int `json:"goodsNum" form:"goodsNum" gorm:"default:0;column:goods_num;comment:;"`  //商品库存 
      GoodsVender  string `json:"goodsVender" form:"goodsVender" gorm:"column:goods_vender;comment:;" binding:"required"`  //商品供货商 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 商品信息 My_goods自定义表名 my_goods
func (My_goods) TableName() string {
  return "my_goods"
}

