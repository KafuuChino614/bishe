package mySys

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mySys"
    mySysReq "github.com/flipped-aurora/gin-vue-admin/server/model/mySys/request"
    "gorm.io/gorm"
)

type CustomerService struct {
}

// CreateCustomer 创建顾客记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerService *CustomerService) CreateCustomer(customer *mySys.Customer) (err error) {
	err = global.GVA_DB.Create(customer).Error
	return err
}

// DeleteCustomer 删除顾客记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerService *CustomerService)DeleteCustomer(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.Customer{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&mySys.Customer{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCustomerByIds 批量删除顾客记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerService *CustomerService)DeleteCustomerByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&mySys.Customer{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&mySys.Customer{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCustomer 更新顾客记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerService *CustomerService)UpdateCustomer(customer mySys.Customer) (err error) {
	err = global.GVA_DB.Model(&mySys.Customer{}).Where("id = ?",customer.ID).Updates(&customer).Error
	return err
}

// GetCustomer 根据ID获取顾客记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerService *CustomerService)GetCustomer(ID string) (customer mySys.Customer, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&customer).Error
	return
}

// GetCustomerInfoList 分页获取顾客记录
// Author [piexlmax](https://github.com/piexlmax)
func (customerService *CustomerService)GetCustomerInfoList(info mySysReq.CustomerSearch) (list []mySys.Customer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mySys.Customer{})
    var customers []mySys.Customer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&customers).Error
	return  customers, total, err
}
