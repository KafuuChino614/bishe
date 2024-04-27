package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
    mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
    "gorm.io/gorm"
)

type WareHouseService struct {
}

// CreateWareHouse 创建仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseService *WareHouseService) CreateWareHouse(wareHouse *mySys.WareHouse) (err error) {
	err = global.GVA_DB.Create(wareHouse).Error
	return err
}

// DeleteWareHouse 删除仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseService *WareHouseService)DeleteWareHouse(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.WareHouse{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&mySys.WareHouse{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWareHouseByIds 批量删除仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseService *WareHouseService)DeleteWareHouseByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.WareHouse{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&mySys.WareHouse{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWareHouse 更新仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseService *WareHouseService)UpdateWareHouse(wareHouse mySys.WareHouse) (err error) {
	err = global.GVA_DB.Model(&mySys.WareHouse{}).Where("id = ?",wareHouse.ID).Updates(&wareHouse).Error
	return err
}

// GetWareHouse 根据ID获取仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseService *WareHouseService)GetWareHouse(ID string) (wareHouse mySys.WareHouse, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wareHouse).Error
	return
}

// GetWareHouseInfoList 分页获取仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (wareHouseService *WareHouseService)GetWareHouseInfoList(info mySysReq.WareHouseSearch) (list []mySys.WareHouse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mySys.WareHouse{})
    var wareHouses []mySys.WareHouse
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.WareHouseName != "" {
        db = db.Where("ware_house_name = ?",info.WareHouseName)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&wareHouses).Error
	return  wareHouses, total, err
}
