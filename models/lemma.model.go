package models

import (
	"rn-game/database"
	"rn-game/schemas"
	"rn-game/utils"
	"github.com/jinzhu/gorm"
)

type LemmaModel struct {
	BaseModel
	database   *gorm.DB
}

// 获取指定数目的词条
func (lm LemmaModel) GetRandomLemmas(number int) ([]schemas.Lemma, error) {
	var lemmas []schemas.Lemma

	count, error := lm.GetLemmasCount()

	if error != nil {
		return lemmas, error
	}

	var randomNumbers []int

	for index := 0; index < number; index++ {
		randomNumbers = append(randomNumbers, utils.GetRandomInt(1, count))
	}

	if error := lm.database.Table("cb_lemmas").Where("id in (?)", randomNumbers).Find(&lemmas).Error; error != nil {
		return lemmas, error
	}
	return lemmas, nil
}

// 获取词条总数目
func (lm LemmaModel) GetLemmasCount() (int, error) {
	var count int = 0

	if error := lm.database.Table("cb_lemmas").Count(&count).Error; error != nil {
		return count, error
	}
	return count, nil
}

// 根据 id 获取指定的词条列表
func (lm LemmaModel) GetLemmasByIds(ids []int) ([]schemas.Lemma, error) {
	var lemmas []schemas.Lemma

	if error := lm.database.Table("cb_lemmas").Where("id in (?)", ids).Find(&lemmas).Error; error != nil {
		return lemmas, error
	}
	return lemmas, nil
}

func NewLemmaModel() LemmaModel {
	return LemmaModel {
		database: database.GetDatabaseHandle(),
	}
}
