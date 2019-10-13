package router

import (
	"rn-game/controllers"
	"github.com/gin-gonic/gin"
)

var indexController controller.IndexController = controller.IndexController{}

var indexRoutes []route = []route {
	route {
		Path: "/index",
		Method: "GET",
		Handles: []gin.HandlerFunc { indexController.Index },
	},
	route {
		Path: "/index/data",
		Method: "GET",
		Handles: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.GetIndexData },
	},
}
