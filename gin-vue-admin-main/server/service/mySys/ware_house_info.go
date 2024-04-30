package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
	mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
	"gorm.io/gorm"
)

type WareHouseInfoService struct {
}

var goodsService *My_goodsService

// CreateWareHouseInfo 创建仓库信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseInfoService *WareHouseInfoService) CreateWareHouseInfo(wareHouseInfo *mySys.WareHouseInfo) (err error) {
	//根据商品ID获取商品信息
	var goods mySys.My_goods
	global.GVA_DB.Where("id = ?", *wareHouseInfo.GoodsID).First(&goods)
	//先根据id查看商品是否存在，存在看仓库id，仓库id一样则将num值++，否则新建
	var iwareHouseInfos []mySys.WareHouseInfo
	global.GVA_DB.Where("goods_i_d=?", *wareHouseInfo.GoodsID).Find(&iwareHouseInfos)
	//查不到 则新建
	if len(iwareHouseInfos) == 0 {
		err = global.GVA_DB.Create(wareHouseInfo).Error
		//更新商品中的数量值
		*goods.GoodsNum -= *wareHouseInfo.Num
		goodsService.UpdateMy_goods(goods)

		return err
	}
	//查到商品,查看仓库ID
	for i := 0; i < len(iwareHouseInfos); i++ {
		if *iwareHouseInfos[i].WareHouseID == *wareHouseInfo.WareHouseID {
			//查到仓库ID,修改商品数量
			*iwareHouseInfos[i].Num += *wareHouseInfo.Num
			wareHouseInfoService.UpdateWareHouseInfo(iwareHouseInfos[i])
			//更新商品中的数量值
			*goods.GoodsNum -= *wareHouseInfo.Num
			goodsService.UpdateMy_goods(goods)
			return err
		}
	}
	//查到商品，没查到仓库,新建
	err = global.GVA_DB.Create(wareHouseInfo).Error
	//更新商品中的数量值
	*goods.GoodsNum -= *wareHouseInfo.Num
	goodsService.UpdateMy_goods(goods)
	return err
}

// DeleteWareHouseInfo 删除仓库信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseInfoService *WareHouseInfoService) DeleteWareHouseInfo(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&mySys.WareHouseInfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&mySys.WareHouseInfo{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWareHouseInfoByIds 批量删除仓库信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseInfoService *WareHouseInfoService) DeleteWareHouseInfoByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&mySys.WareHouseInfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&mySys.WareHouseInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWareHouseInfo 更新仓库信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseInfoService *WareHouseInfoService) UpdateWareHouseInfo(wareHouseInfo mySys.WareHouseInfo) (err error) {
	err = global.GVA_DB.Model(&mySys.WareHouseInfo{}).Where("id = ?", wareHouseInfo.ID).Updates(&wareHouseInfo).Error
	return err
}

// GetWareHouseInfo 根据ID获取仓库信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseInfoService *WareHouseInfoService) GetWareHouseInfo(ID string) (wareHouseInfo mySys.WareHouseInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wareHouseInfo).Error
	return
}

// GetWareHouseInfoInfoList 分页获取仓库信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseInfoService *WareHouseInfoService) GetWareHouseInfoInfoList(info mySysReq.WareHouseInfoSearch) (list []mySys.WareHouseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mySys.WareHouseInfo{})
	var wareHouseInfos []mySys.WareHouseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.WareHouseName != "" {
		db = db.Where("ware_house_name = ?", info.WareHouseName)
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name = ?", info.GoodsName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wareHouseInfos).Error
	return wareHouseInfos, total, err
}
