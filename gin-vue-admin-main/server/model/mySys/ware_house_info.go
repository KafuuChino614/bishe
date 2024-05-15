// 自动生成模板WareHouseInfo
package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 仓库信息 结构体  WareHouseInfo
type WareHouseInfo struct {
	global.GVA_MODEL
	WareHouseID     *int       `json:"wareHouseID" form:"wareHouseID" gorm:"column:ware_house_i_d;comment:仓库ID;" binding:"required"`      //仓库ID
	GoodsID         *int       `json:"goodsID" form:"goodsID" gorm:"column:goods_i_d;comment:商品ID;" binding:"required"`                   //商品ID
	Num             *int       `json:"num" form:"num" gorm:"default:0;column:num;comment:商品数量（仓库里的）;" binding:"required"`                 //商品数量
	Price           *float64   `json:"price" form:"price" gorm:"default:0;column:price;comment:;" binding:"required"`                     //商品单价
	WareHouseName   string     `json:"wareHouseName" form:"wareHouseName" gorm:"column:ware_house_name;comment:仓库名字;" binding:"required"` //仓库名字
	GoodsName       string     `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名字;" binding:"required"`              //商品名字
	GoodsType       string     `json:"goodsType" form:"goodsType" gorm:"column:goods_type;comment:商品类型;" binding:"required"`              //商品类型
	GoodsUnit       string     `json:"goodsUnit" form:"goodsUnit" gorm:"column:goods_unit;comment:计量单位;"`                                 //计量单位
	ManufactureData *time.Time `json:"manufactureData" form:"manufactureData" gorm:"column:manufacture_data;comment:;"`                   //生产日期
	ExpirationDate  *time.Time `json:"expirationDate" form:"expirationDate" gorm:"column:expiration_date;comment:过期日期;"`                  //过期日期
	CreatedBy       uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 仓库信息 WareHouseInfo自定义表名 wareHouseInfo
func (WareHouseInfo) TableName() string {
	return "wareHouseInfo"
}
