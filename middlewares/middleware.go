package middleware

import "github.com/gin-gonic/gin"

var globalMiddleware []gin.HandlerFunc

func init() {
	globalMiddleware = append(globalMiddleware, CorsMiddleware)
}

func RegisterMiddleware(engine *gin.Engine) {
	engine.Use(globalMiddleware...)
}
