package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
	mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
	"gorm.io/gorm"
)

type My_goodsTypeService struct {
}

// CreateMy_goodsType 创建商品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsTypeService *My_goodsTypeService) CreateMy_goodsType(my_goodsType *mySys.My_goodsType) (err error) {
	err = global.GVA_DB.Create(my_goodsType).Error
	return err
}

// DeleteMy_goodsType 删除商品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsTypeService *My_goodsTypeService) DeleteMy_goodsType(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&mySys.My_goodsType{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&mySys.My_goodsType{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMy_goodsTypeByIds 批量删除商品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsTypeService *My_goodsTypeService) DeleteMy_goodsTypeByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&mySys.My_goodsType{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&mySys.My_goodsType{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMy_goodsType 更新商品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsTypeService *My_goodsTypeService) UpdateMy_goodsType(my_goodsType mySys.My_goodsType) (err error) {
	err = global.GVA_DB.Model(&mySys.My_goodsType{}).Where("id = ?", my_goodsType.ID).Updates(&my_goodsType).Error
	return err
}

// GetMy_goodsType 根据ID获取商品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsTypeService *My_goodsTypeService) GetMy_goodsType(ID string) (my_goodsType mySys.My_goodsType, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&my_goodsType).Error
	return
}

// GetMy_goodsTypeInfoList 分页获取商品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_goodsTypeService *My_goodsTypeService) GetMy_goodsTypeInfoList(info mySysReq.My_goodsTypeSearch) (list []mySys.My_goodsType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mySys.My_goodsType{})
	var my_goodsTypes []mySys.My_goodsType
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TypeName != "" {
		db = db.Where("type_name = ?", info.TypeName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&my_goodsTypes).Error
	return my_goodsTypes, total, err
}
