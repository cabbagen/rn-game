package controller

import (
	"strconv"
	"rn-game/models"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	BaseController
}


// 获取游戏分类
func (gc GameController) GetAllGameCategories(c *gin.Context) {
	categories, error := models.NewCategoryModel().GetAllCategories()

	if error != nil {
		gc.HandleErrorResponse(c, error)
		return
	}
	gc.HandleSuccessResponse(c, categories)
}

// 根据分类获取游戏列表
func (gc GameController) GetGamesByCategoryId(c *gin.Context) {
	var params struct {
		CategoryId     int    `form:"categoryId" json:"categoryId"`
		PageNo         int    `form:"pageNo" json:"pageNo"`
		PageSize       int    `form:"pageSize" json:"pageSize"`
	}

	if error := c.BindQuery(&params); error != nil {
		gc.HandleErrorResponse(c, error)
		return
	}

	games, total, error := models.NewGameModel().GetGamesByCategoryId(params.CategoryId, params.PageNo, params.PageSize)

	if error != nil {
		gc.HandleErrorResponse(c, error)
		return
	}

	gc.HandleSuccessResponse(c, map[string]interface{} { "games": games, "total": total })
}

// 获取游戏详情
func (gc GameController) GetGameById(c *gin.Context) {
	gameId, error := strconv.Atoi(c.Param("gameId"))

	if error != nil {
		gc.HandleErrorResponse(c, error)
		return
	}

	gameInfo, error := models.NewGameModel().GetGameDetailInfo(gameId)

	if error != nil {
		gc.HandleErrorResponse(c, error)
		return
	}

	gc.HandleSuccessResponse(c, gameInfo)
}

// 获取相似游戏
func (gc GameController) GetLikeGames(c *gin.Context) {
	var params struct {
		SearchKey      string   `form:"searchKey" json:"searchKey"`
		PageNo         int      `form:"pageNo" json:"pageNo"`
		PageSize       int      `form:"pageSize" json:"pageSize"`
	}

	if error := c.BindQuery(&params); error != nil {
		gc.HandleErrorResponse(c, error)
		return
	}

	games, total, error := models.NewGameModel().GetRelatedGames(params.SearchKey, params.PageNo, params.PageSize)

	if error != nil {
		gc.HandleErrorResponse(c, error)
		return
	}

	gc.HandleSuccessResponse(c, map[string]interface{} { "total": total, "games": games })
}
