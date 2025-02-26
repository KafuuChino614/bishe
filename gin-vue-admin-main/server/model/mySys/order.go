// 自动生成模板Order
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 订单 结构体  Order
type Order struct {
 global.GVA_MODEL 
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:18位的订单编号;"`  //订单编号 
      CustomerName  string `json:"customer" form:"customer" gorm:"column:customer;comment:客户姓名;" binding:"required"`  //客户姓名 
      WareHouseName  string `json:"wareHouseName" form:"wareHouseName" gorm:"column:ware_house_name;comment:wareHouseName;" binding:"required"`  //仓库名 
      GoodsName  string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名;" binding:"required"`  //商品名 
      GoodsType  string `json:"goodsType" form:"goodsType" gorm:"column:goods_type;comment:商品类型;" binding:"required"`  //商品类型 
      GoodsUnit  string `json:"goodsUnit" form:"goodsUnit" gorm:"column:goods_unit;comment:商品类型;" binding:"required"`  //计量单位 
      GoodsNum  *int `json:"goodsNum" form:"goodsNum" gorm:"column:goods_num;comment:商品数量;" binding:"required"`  //商品数量 
      GoodsPrice  *float64 `json:"goodsPrice" form:"goodsPrice" gorm:"column:goods_price;comment:商品价格;" binding:"required"`  //商品价格 
      Discount  string `json:"discount" form:"discount" gorm:"column:discount;comment:折扣;" binding:"required"`  //折扣 
      AllPrice  *float64 `json:"allPrice" form:"allPrice" gorm:"column:all_price;comment:总价;"`  //总价 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 订单 Order自定义表名 order
func (Order) TableName() string {
  return "order"
}

