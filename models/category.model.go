package models

import (
	"rn-game/schemas"
	"rn-game/database"
	"github.com/jinzhu/gorm"
)

type CategoryModel struct {
	BaseModel
	database   *gorm.DB
}

/**
 * 获取游戏分类列表
 */
func (cm CategoryModel) GetAllCategories() ([]schemas.Category, error) {
	var categories []schemas.Category

	if error := cm.database.Table("cb_categories").Find(&categories).Error; error != nil {
		return categories, error
	}
	return categories, nil
}

func NewCategoryModel() CategoryModel {
	return CategoryModel {
		database: database.GetDatabaseHandle(),
	}
}
