package router

import (
	"rn-game/controllers"
	"github.com/gin-gonic/gin"
)

var gameController controller.GameController = controller.GameController{}

var gameRoutes []route = []route {
	route {
		Path: "/game/categories",
		Method: "GET",
		Handles: []gin.HandlerFunc { gameController.GetAllGameCategories },
	},
	route {
		Path: "/game/games",
		Method: "GET",
		Handles: []gin.HandlerFunc { gameController.GetGamesByCategoryId },
	},
	route {
		Path: "/game/game/:gameId",
		Method: "GET",
		Handles: []gin.HandlerFunc { tokenMiddleware.Handle, gameController.GetGameById },
	},
	route {
		Path: "/game/games/like",
		Method: "GET",
		Handles: []gin.HandlerFunc { gameController.GetLikeGames },
	},
}
