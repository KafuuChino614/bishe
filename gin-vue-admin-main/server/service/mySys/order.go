package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
    mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
    "gorm.io/gorm"
)

type OrderService struct {
}

// CreateOrder 创建订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) CreateOrder(order *mySys.Order) (err error) {
	err = global.GVA_DB.Create(order).Error
	return err
}

// DeleteOrder 删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService)DeleteOrder(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.Order{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&mySys.Order{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOrderByIds 批量删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService)DeleteOrderByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.Order{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&mySys.Order{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOrder 更新订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService)UpdateOrder(order mySys.Order) (err error) {
	err = global.GVA_DB.Model(&mySys.Order{}).Where("id = ?",order.ID).Updates(&order).Error
	return err
}

// GetOrder 根据ID获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService)GetOrder(ID string) (order mySys.Order, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&order).Error
	return
}

// GetOrderInfoList 分页获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService)GetOrderInfoList(info mySysReq.OrderSearch) (list []mySys.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mySys.Order{})
    var orders []mySys.Order
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Uuid != "" {
        db = db.Where("uuid = ?",info.Uuid)
    }
    if info.CustomerName != "" {
        db = db.Where("customer = ?",info.CustomerName)
    }
    if info.WareHouseName != "" {
        db = db.Where("ware_house_name = ?",info.WareHouseName)
    }
    if info.GoodsName != "" {
        db = db.Where("goods_name = ?",info.GoodsName)
    }
    if info.GoodsType != "" {
        db = db.Where("goods_type = ?",info.GoodsType)
    }
    if info.GoodsUnit != "" {
        db = db.Where("goods_unit = ?",info.GoodsUnit)
    }
    if info.GoodsPrice != nil {
        db = db.Where("goods_price = ?",info.GoodsPrice)
    }
    if info.Discount != "" {
        db = db.Where("discount = ?",info.Discount)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&orders).Error
	return  orders, total, err
}
