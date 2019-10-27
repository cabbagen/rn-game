package controller

import (
	"github.com/gin-gonic/gin"
	"rn-game/schemas"
	"rn-game/models"
)

type LemmaController struct {
	BaseController
}

func (lc LemmaController) RenderLemmaGame(c *gin.Context) {
	lemmas, error := lc.getRandomLemmas(2)

	if error != nil {
		lc.HandleErrorResponse(c, error)
		return
	}
	c.HTML(200, "lemma.html", gin.H {
		"data": gin.H {
			"lemmas": lemmas,
		},
	})
}

// 获取随机词条
func (lc LemmaController) getRandomLemmas(number int) ([]schemas.Lemma, error) {
	lemmas, error := models.NewLemmaModel().GetRandomLemmas(number)

	if error != nil {
		return nil, error
	}
	return lemmas, nil
}

