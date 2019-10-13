package models

import (
	"rn-game/schemas"
	"rn-game/database"
	"github.com/jinzhu/gorm"
)

type GameModel struct {
	BaseModel
	database   *gorm.DB
}

/**
 * 按分类查找游戏列表
 */
func (gm GameModel) GetGamesByCategoryId(categoryId, pageNo, pageSize int) ([]schemas.Game, int, error) {
	var total int	
	var games []schemas.Game

	error := gm.database.Table("cb_games").Where("category_id = ? and status = 1", categoryId).Count(&total).Offset(pageNo * pageSize).Limit(pageSize).Find(&games).Error

	if error != nil {
		return games, total, error
	}

	return games, total, nil
}

/**
 * 按名称模糊查询游戏
 */
func (gm GameModel) GetRelatedGames(search string, pageNo, pageSize int) ([]schemas.Game, int, error) {
	var total int
	var games []schemas.Game

	error := gm.database.Table("cb_games").Where("name like ? and status = 1", "%" + search + "%").Count(&total).Offset(pageSize * pageNo).Limit(pageSize).Find(&games).Error

	if error != nil {
		return games, total, error
	}
	return games, total, nil
}

/**
 * 获取游戏详情
 */
func (gm GameModel) GetGameDetailInfo(gameId int) (schemas.GameDetail, error) {
	var gameDetail schemas.GameDetail

	error := gm.database.
		Table("cb_games").
		Select("cb_games.*, cb_categories.name as categoryName").
		Joins("inner join cb_categories on cb_games.category_id = cb_categories.id").
		Where("cb_games.id = ?", gameId).
		Scan(&gameDetail).
		Error

	if error != nil {
		return gameDetail, error
	}
	return gameDetail, nil
}

// 根据 Id 获取指定的游戏
func (gm GameModel) GetGamesByIds(gameIds []int) ([]schemas.Game, error) {
	var games []schemas.Game

	if error := gm.database.Table("cb_games").Where("id in (?)", gameIds).Find(&games).Error; error != nil {
		return games, error
	}
	return games, nil
}

func NewGameModel() GameModel {
	return GameModel {
		database: database.GetDatabaseHandle(),
	}
}
