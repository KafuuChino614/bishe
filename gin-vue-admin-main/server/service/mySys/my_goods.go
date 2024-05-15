package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
	mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
	"gorm.io/gorm"
)

type My_goodsService struct {
}

// CreateMy_goods 创建商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsService *My_goodsService) CreateMy_goods(my_goods *mySys.My_goods) (err error) {
	err = global.GVA_DB.Create(my_goods).Error
	return err
}

// DeleteMy_goods 删除商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsService *My_goodsService) DeleteMy_goods(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&mySys.My_goods{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&mySys.My_goods{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMy_goodsByIds 批量删除商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsService *My_goodsService) DeleteMy_goodsByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&mySys.My_goods{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&mySys.My_goods{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMy_goods 更新商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsService *My_goodsService) UpdateMy_goods(my_goods mySys.My_goods) (err error) {
	err = global.GVA_DB.Model(&mySys.My_goods{}).Where("id = ?", my_goods.ID).Updates(&my_goods).Error
	//更新商品时需要同时更新wareHouseInfo表里对应商品的信息
	goodsID := my_goods.ID
	goodsName := my_goods.GoodsName
	goodsType := my_goods.GoodsType
	goodsUnit := my_goods.GoodsUnit
	manufactureData := my_goods.ManufactureData
	expirationDate := my_goods.ExpirationDate
	//根据goodsID查询wareHouseInfo表中对应的商品，并修改
	// 使用原生 SQL 语句更新 wareHouseInfo 表中对应商品的信息
	sql := "UPDATE warehouseinfo SET goods_name = ?, goods_type = ?, goods_unit = ?,manufacture_data = ?,expiration_date = ? WHERE goods_i_d = ?"
	err = global.GVA_DB.Exec(sql, goodsName, goodsType, goodsUnit, manufactureData, expirationDate, goodsID).Error
	if err != nil {
		return err
	}
	return err
}

// GetMy_goods 根据ID获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsService *My_goodsService) GetMy_goods(ID string) (my_goods mySys.My_goods, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&my_goods).Error
	return
}

// GetMy_goodsInfoList 分页获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsService *My_goodsService) GetMy_goodsInfoList(info mySysReq.My_goodsSearch) (list []mySys.My_goods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mySys.My_goods{})
	var my_goodss []mySys.My_goods
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name = ?", info.GoodsName)
	}
	if info.GoodsType != "" {
		db = db.Where("goods_type = ?", info.GoodsType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&my_goodss).Error
	return my_goodss, total, err
}
