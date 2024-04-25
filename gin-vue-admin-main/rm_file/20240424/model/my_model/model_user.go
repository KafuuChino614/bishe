// 自动生成模板Model_user
package my_model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 用户信息 结构体  Model_user
type Model_user struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:用户名;" binding:"required"`  //用户名 
      Password  string `json:"password" form:"password" gorm:"column:password;comment:用户密码;" binding:"required"`  //密码 
      Auth  *int `json:"auth" form:"auth" gorm:"default:1;column:auth;comment:用户权限(0为管理员,1为普通用户);" binding:"required"`  //权限 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 用户信息 Model_user自定义表名 model_user
func (Model_user) TableName() string {
  return "model_user"
}

