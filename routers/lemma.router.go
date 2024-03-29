package router

import (
	"rn-game/controllers"
	"github.com/gin-gonic/gin"
)

var lemmaController controller.LemmaController = controller.LemmaController{}

var lemmaRoutes []route = []route {
	route {
		Path: "/lemmas/game",
		Method: "GET",
		Handles: []gin.HandlerFunc { lemmaController.RenderLemmaGame },
	},
}
