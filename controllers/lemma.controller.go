package controller

import (
	"github.com/gin-gonic/gin"
	"rn-game/models"
	"rn-game/utils"
	"strconv"
)

type LemmaController struct {
	BaseController
}


// 获取随机词条
func(lc LemmaController) GetRandomLemmas(c * gin.Context) {
	number, error := strconv.Atoi(c.Param("number"))

	if error != nil {
		lc.HandleErrorResponse(c, error)
		return
	}

	lemmas, error := models.NewLemmaModel().GetRandomLemmas(number)

	if error != nil {
		lc.HandleErrorResponse(c, error)
		return
	}

	lc.HandleSuccessResponse(c, lemmas)
}

// 校验随机词条
func (lc LemmaController) HandleValidateLemmas(c *gin.Context) {
	var json struct {
		Ids       []int     `json:"ids"`
		Lemmas    []string  `json:"lemmas"`
	}
	if error := c.ShouldBindJSON(&json); error != nil {
		lc.HandleErrorResponse(c, error)
		return
	}

	lemmas, error := models.NewLemmaModel().GetLemmasByIds(json.Ids)

	if error != nil {
		lc.HandleErrorResponse(c, error)
		return
	}

	var isValidated bool = true

	for _, lemma := range lemmas {
		if !utils.IncludeInStringList(json.Lemmas, lemma.Name) {
			isValidated = false
		}
	}

	lc.HandleSuccessResponse(c, isValidated)
}