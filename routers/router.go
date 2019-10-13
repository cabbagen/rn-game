package router

import (
	"rn-game/middlewares"
	"github.com/gin-gonic/gin"
)

type route struct {
	Path     string
	Method   string
	Handles  []gin.HandlerFunc
}

var routes []route
var tokenMiddleware middleware.TokenMiddleware = middleware.NewTokenMiddleware()

func init() {
	routes = append(routes, indexRoutes...)
	routes = append(routes, gameRoutes...)
	routes = append(routes, lemmaRoutes...)
	routes = append(routes, userRoutes...)
}

func RegisterRouters(engine *gin.Engine) {
	for _, route := range routes {
		engine.Handle(route.Method, route.Path, route.Handles...)
	}
	engine.Static("/static", "./resources")
}
