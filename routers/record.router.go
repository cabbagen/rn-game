package router

import (
	"rn-game/controllers"
	"github.com/gin-gonic/gin"
)

var recordController controller.RecordController = controller.RecordController{}

var recordRoutes []route = []route {
	route {
		Path: "/record/records",
		Method: "GET",
		Handles: []gin.HandlerFunc { recordController.GetUserGameRecords },
	},
}