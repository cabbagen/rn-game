package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"rn-game/schemas"
	"encoding/json"
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

func (bc BaseController) GetAuthUserId(c *gin.Context) (int, error) {
	var userInfo schemas.User

	userInfoString, isExist := c.Get("userInfo")

	if !isExist {
		return userInfo.ID, errors.New("用户不存在")
	}

	if error := json.Unmarshal([]byte(userInfoString.(string)), &userInfo); error != nil {
		return userInfo.ID, error
	}

	return userInfo.ID, nil
}

