package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
    mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
    "gorm.io/gorm"
)

type My_vendorService struct {
}

// CreateMy_vendor 创建供货商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_vendorService *My_vendorService) CreateMy_vendor(my_vendor *mySys.My_vendor) (err error) {
	err = global.GVA_DB.Create(my_vendor).Error
	return err
}

// DeleteMy_vendor 删除供货商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_vendorService *My_vendorService)DeleteMy_vendor(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.My_vendor{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&mySys.My_vendor{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMy_vendorByIds 批量删除供货商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_vendorService *My_vendorService)DeleteMy_vendorByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.My_vendor{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&mySys.My_vendor{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMy_vendor 更新供货商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_vendorService *My_vendorService)UpdateMy_vendor(my_vendor mySys.My_vendor) (err error) {
	err = global.GVA_DB.Model(&mySys.My_vendor{}).Where("id = ?",my_vendor.ID).Updates(&my_vendor).Error
	return err
}

// GetMy_vendor 根据ID获取供货商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_vendorService *My_vendorService)GetMy_vendor(ID string) (my_vendor mySys.My_vendor, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&my_vendor).Error
	return
}

// GetMy_vendorInfoList 分页获取供货商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (my_vendorService *My_vendorService)GetMy_vendorInfoList(info mySysReq.My_vendorSearch) (list []mySys.My_vendor, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mySys.My_vendor{})
    var my_vendors []mySys.My_vendor
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.VenderName != "" {
        db = db.Where("vender_name = ?",info.VenderName)
    }
    if info.VenderAddr != "" {
        db = db.Where("vender_addr = ?",info.VenderAddr)
    }
    if info.Header != "" {
        db = db.Where("header = ?",info.Header)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&my_vendors).Error
	return  my_vendors, total, err
}
