package my_model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/my_model"
    my_modelReq "github.com/flipped-aurora/gin-vue-admin/server/model/my_model/request"
    "gorm.io/gorm"
)

type Model_userService struct {
}

// CreateModel_user 创建用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (model_userService *Model_userService) CreateModel_user(model_user *my_model.Model_user) (err error) {
	err = global.GVA_DB.Create(model_user).Error
	return err
}

// DeleteModel_user 删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (model_userService *Model_userService)DeleteModel_user(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&my_model.Model_user{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&my_model.Model_user{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteModel_userByIds 批量删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (model_userService *Model_userService)DeleteModel_userByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&my_model.Model_user{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&my_model.Model_user{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateModel_user 更新用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (model_userService *Model_userService)UpdateModel_user(model_user my_model.Model_user) (err error) {
	err = global.GVA_DB.Model(&my_model.Model_user{}).Where("id = ?",model_user.ID).Updates(&model_user).Error
	return err
}

// GetModel_user 根据ID获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (model_userService *Model_userService)GetModel_user(ID string) (model_user my_model.Model_user, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&model_user).Error
	return
}

// GetModel_userInfoList 分页获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (model_userService *Model_userService)GetModel_userInfoList(info my_modelReq.Model_userSearch) (list []my_model.Model_user, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&my_model.Model_user{})
    var model_users []my_model.Model_user
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
    if info.Password != "" {
        db = db.Where("password = ?",info.Password)
    }
    if info.Auth != nil {
        db = db.Where("auth = ?",info.Auth)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&model_users).Error
	return  model_users, total, err
}
