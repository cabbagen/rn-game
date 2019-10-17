package view

import "github.com/gin-gonic/gin"

var pattern string = "./views/*.html"

func RegisterTemplate(engine *gin.Engine) {
	engine.Delims("[[", "]]").LoadHTMLGlob(pattern)
}
