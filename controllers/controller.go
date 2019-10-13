package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	
}

func (bc BaseController) DownloadFile(c *gin.Context, filePath, filename string) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.File(filePath)
}

func (bc BaseController) PreviewFile(c *gin.Context, filePath, filename string) {
	c.Header("Content-Disposition", fmt.Sprintf("filename=%s", filename))
	c.File(filePath)
}

func (bc BaseController) HandleErrorResponse(c *gin.Context, e error) {
	c.JSON(200, gin.H{ "status": 500, "data": nil, "msg": e.Error() })
}

func (bc BaseController) HandleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{ "status": 200, "data": data, "msg": nil })
}