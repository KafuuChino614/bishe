package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
    mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
    "gorm.io/gorm"
)

type My_goodsUnitService struct {
}

// CreateMy_goodsUnit 创建计量单位记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsUnitService *My_goodsUnitService) CreateMy_goodsUnit(my_goodsUnit *mySys.My_goodsUnit) (err error) {
	err = global.GVA_DB.Create(my_goodsUnit).Error
	return err
}

// DeleteMy_goodsUnit 删除计量单位记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsUnitService *My_goodsUnitService)DeleteMy_goodsUnit(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.My_goodsUnit{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&mySys.My_goodsUnit{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMy_goodsUnitByIds 批量删除计量单位记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsUnitService *My_goodsUnitService)DeleteMy_goodsUnitByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.My_goodsUnit{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&mySys.My_goodsUnit{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMy_goodsUnit 更新计量单位记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsUnitService *My_goodsUnitService)UpdateMy_goodsUnit(my_goodsUnit mySys.My_goodsUnit) (err error) {
	err = global.GVA_DB.Model(&mySys.My_goodsUnit{}).Where("id = ?",my_goodsUnit.ID).Updates(&my_goodsUnit).Error
	return err
}

// GetMy_goodsUnit 根据ID获取计量单位记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsUnitService *My_goodsUnitService)GetMy_goodsUnit(ID string) (my_goodsUnit mySys.My_goodsUnit, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&my_goodsUnit).Error
	return
}

// GetMy_goodsUnitInfoList 分页获取计量单位记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsUnitService *My_goodsUnitService)GetMy_goodsUnitInfoList(info mySysReq.My_goodsUnitSearch) (list []mySys.My_goodsUnit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mySys.My_goodsUnit{})
    var my_goodsUnits []mySys.My_goodsUnit
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&my_goodsUnits).Error
	return  my_goodsUnits, total, err
}
