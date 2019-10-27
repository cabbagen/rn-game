package models

import (
	"rn-game/database"
	"rn-game/schemas"
	"github.com/jinzhu/gorm"
)

type RecordModel struct {
	BaseModel
	database   *gorm.DB
}


// 添加游戏记录
func(rm RecordModel) AddGameRecord(userId, gameId int) (bool, error) {
	record := schemas.Record {
		UserId: userId,
		GameId: gameId,
	}

	if error := rm.database.Table("cb_records").Create(&record).Error; error != nil {
		return false, error
	}
	return true, nil
}

// 查询用户游戏记录
func(rm RecordModel) GetUserRecords(userId int, startTime, endTime string, pageNo, pageSize int) ([]schemas.RecordDetail, int, error) {
	var total int
	var records []schemas.RecordDetail
	var times []string = []string{ "1970-01-01", "3010-01-01" }

	if startTime != "" {
		times[0] = startTime
	}

	if endTime != "" {
		times[1] = endTime
	}


	error := rm.database.
		Table("cb_records").
		Select("cb_records.id as id, username, nickname, gender, avatar, cb_games.img as gameImg, cb_games.name as gameName, cb_games.play_link as gamePlayLink, cb_categories.name as gameCategoryName").
		Where("user_id = ? and cb_records.created_at > ? and cb_records.created_at < ?", userId, times[0], times[1]).
		Joins("inner join cb_games on cb_records.game_id = cb_games.id").
		Joins("inner join cb_users on cb_records.user_id = cb_users.id").
		Joins("inner join cb_categories on cb_categories.id = cb_games.category_id").
		Count(&total).
		Offset(pageNo * pageSize).
		Limit(pageSize).
		Scan(&records).
		Error

	if error != nil {
		return records, total, error
	}
	return records, total, nil
}


func NewRecordModel() RecordModel {
	return RecordModel {
		database: database.GetDatabaseHandle(),
	}
}
