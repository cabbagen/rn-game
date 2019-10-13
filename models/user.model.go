package models

import (
	"rn-game/schemas"
	"rn-game/database"
	"github.com/jinzhu/gorm"
	"rn-game/utils"
)

type UserModel struct {
	BaseModel
	database   *gorm.DB
}

// 通过 id 获取用户信息
func (um UserModel) GetUserInfoById(userId int) (schemas.User, error) {
	var user schemas.User

	if error := um.database.First(&user, userId).Error; error != nil {
		return user, error
	}

	return user, nil
}

// 检验用户是否存在
func (um UserModel) CheckIsExistUser(username, password string) (schemas.User, bool) {
	var user schemas.User

	if error := um.database.Where("username = ? and password = ?", username, utils.MD5(password)).First(&user).Error; error != nil {
		return user, false
	}
	return user, true
}

func NewUserModel() UserModel {
	return UserModel {
		database: database.GetDatabaseHandle(),
	}
}
