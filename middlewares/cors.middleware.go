package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var headers map[string]string = map[string]string {
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Methods": "GET, POST, HEAD, DELETE, PUT, OPTIONS",
	"Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type",
}

func CorsMiddleware(c *gin.Context) {
	for key, value := range headers {
		c.Header(key, value)
	}

	if c.Request.Method == "OPTIONS" {
		c.String(http.StatusOK, "true")
	}
	c.Next()
}
