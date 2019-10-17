package router

import (
	"rn-game/controllers"
	"github.com/gin-gonic/gin"
)

var userController controller.UserController = controller.UserController{}

var userRoutes []route = []route {
	route {
		Path: "/user/login",
		Method: "POST",
		Handles: []gin.HandlerFunc { userController.Login },
	},
	route {
		Path: "/user/userInfo",
		Method: "GET",
		Handles: []gin.HandlerFunc { tokenMiddleware.Handle, userController.GetUserInfo },
	},
}
