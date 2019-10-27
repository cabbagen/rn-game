package controller

import (
	"github.com/gin-gonic/gin"
	"rn-game/models"
)

type RecordController struct {
	BaseController
}

func (rc RecordController) GetUserGameRecords(c *gin.Context) {
	var params struct {
		UserId         int       `form:"userId" json:"userId"`
		StartTime      string    `form:"startTime" json:"startTime"`
		EndTime        string    `form:"endTime" json:"endTime"`
		PageNo         int       `form:"pageNo" json:"pageNo"`
		PageSize       int       `form:"pageSize" json:"pageSize"`
	}

	if error := c.BindQuery(&params); error != nil {
		rc.HandleErrorResponse(c, error)
		return
	}

	records, total, error := models.NewRecordModel().GetUserRecords(params.UserId, params.StartTime, params.EndTime, params.PageNo, params.PageSize)

	if error != nil {
		rc.HandleErrorResponse(c, error)
		return
	}

	rc.HandleSuccessResponse(c, gin.H { "records": records, "total": total })
}

