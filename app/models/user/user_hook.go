// Gorm 提供了 BeforeSave 的模型钩子，会在模型创建和更新前被调用，我们利用此机制在入库前对密码做加密：
package user

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (userModel *User) BeforeSave(db *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}

	return
}
