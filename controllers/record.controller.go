package controller

import (
	"github.com/gin-gonic/gin"
	"rn-game/models"
)

type RecordController struct {
	BaseController
}

// 添加用户游戏记录
func (rc RecordController) AddUserGameRecord(c *gin.Context) {
	var params struct {
		GameId         int       `form:"gameId" json:"gameId"`
	}

	if error := c.BindJSON(&params); error != nil {
		rc.HandleErrorResponse(c, error)
		return
	}

	userId, error := rc.GetAuthUserId(c)

	if error != nil {
		rc.HandleErrorResponse(c, error)
		return
	}

	if result, _ := models.NewRecordModel().AddGameRecord(userId, params.GameId); !result {
		rc.HandleErrorResponse(c, error)
		return
	}

	rc.HandleSuccessResponse(c, "ok")
}

// 获取用户游戏记录
func (rc RecordController) GetUserGameRecords(c *gin.Context) {
	var params struct {
		StartTime      string    `form:"startTime" json:"startTime"`
		EndTime        string    `form:"endTime" json:"endTime"`
		PageNo         int       `form:"pageNo" json:"pageNo"`
		PageSize       int       `form:"pageSize" json:"pageSize"`
	}

	if error := c.BindQuery(&params); error != nil {
		rc.HandleErrorResponse(c, error)
		return
	}

	userId, error := rc.GetAuthUserId(c)

	if error != nil {
		rc.HandleErrorResponse(c, error)
		return
	}

	records, total, error := models.NewRecordModel().GetUserRecords(userId, params.StartTime, params.EndTime, params.PageNo, params.PageSize)

	if error != nil {
		rc.HandleErrorResponse(c, error)
		return
	}

	rc.HandleSuccessResponse(c, gin.H { "records": records, "total": total })
}

