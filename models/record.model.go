package models

import (
	"rn-game/database"
	"github.com/jinzhu/gorm"
)

type RecordModel struct {
	BaseModel
	database   *gorm.DB
}

func NewRecordModel() RecordModel {
	return RecordModel {
		database: database.GetDatabaseHandle(),
	}
}
