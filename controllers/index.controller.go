package controller

import (
	"net/http"
	"rn-game/schemas"
	"github.com/gin-gonic/gin"
	"rn-game/models"
	"rn-game/utils"
)

type IndexController struct {
	BaseController
}

// 问候地址
func (ic IndexController) Index (c *gin.Context) {
	c.String(http.StatusOK, "hello index")
}

type IndexCategory struct {
	Title          string                   `json:"title"`
	Games          []schemas.Game           `json:"games"`
}


func (ic IndexController) GetIndexData(c *gin.Context) {
	var indexData struct {
		Banners       []string              `json:"banners"`
		Categories    []IndexCategory       `json:"categories"`
	}

	indexCategories, error := ic.getIndexCategories()

	if error != nil {
		ic.HandleErrorResponse(c, error)
		return
	}

	indexData.Banners = ic.getIndexBanner()
	indexData.Categories = indexCategories

	ic.HandleSuccessResponse(c, indexData)
}

func (ic IndexController) getIndexBanner() []string {
	return []string {
		"https://gw.alicdn.com/tps/i4/TB1PlHsgkT2gK0jSZPcSuwKkpXa.jpg_790x10000Q75.jpg",
		"https://gw.alicdn.com/imgextra/i3/1562021/O1CN01s52R2Q1Qnf0fNJh9I_!!1562021-0-lubanu.jpg_790x10000Q75.jpg",
		"https://gw.alicdn.com/imgextra/i4/1453939/O1CN01xBT2G41ey6fJsagww_!!1453939-0-lubanu.jpg_790x10000Q75.jpg",
		"https://gw.alicdn.com/tps/i4/TB1PlHsgkT2gK0jSZPcSuwKkpXa.jpg_790x10000Q75.jpg",
		"https://gw.alicdn.com/imgextra/i1/30625/O1CN01f94dzV1GUI6GrmG4I_!!30625-0-lubanu.jpg_790x10000Q75.jpg",
	}
}

func (ic IndexController) getIndexCategories() ([]IndexCategory, error) {
	var allGameIds []int
	var indexCategory []IndexCategory

	var categoryInfo map[string][]int = map[string][]int {
		"热门推荐": []int { 29, 37, 45, 60, 72 },
		"经典射击": []int { 264, 284, 287, 241, 240 },
		"午后休闲": []int { 92, 94, 86, 80, 79 },
		"体育飙车": []int { 181, 182, 185, 186, 189 },
	}

	for _, value := range categoryInfo {
		allGameIds = append(allGameIds, value...)
	}

	games, error := models.NewGameModel().GetGamesByIds(allGameIds)

	if error != nil {
		return indexCategory, error
	}

	for key, value := range categoryInfo {
		current := IndexCategory{
			Title: key,
			Games: []schemas.Game {},
		}
		for _, game := range games {
			if utils.IncludeInIntList(value, game.ID) {
				current.Games = append(current.Games, game)
			}
		}

		indexCategory = append(indexCategory, current)
	}

	return indexCategory, nil
}

