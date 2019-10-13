package controller

import (
	"github.com/gin-gonic/gin"
	"rn-game/models"
	"github.com/pkg/errors"
	"encoding/json"
	"rn-game/middlewares"
)

type UserController struct {
	BaseController
}

// 用户登录
func (uc UserController) Login(c *gin.Context) {
	var params struct {
		Username      string    `json:"username"`
		Password      string    `json:"password"`
	}

	if error := c.BindJSON(&params); error != nil {
		uc.HandleErrorResponse(c, error)
	}

	userInfo, isExist := models.NewUserModel().CheckIsExistUser(params.Username, params.Password)

	if !isExist {
		uc.HandleErrorResponse(c, errors.New("用户不存在"))
		return
	}

	userInfoByte, error := json.Marshal(userInfo)

	if error != nil {
		uc.HandleErrorResponse(c, errors.New("用户不存在"))
		return
	}

	tokenString, error := middleware.NewTokenMiddleware().SignToken(string(userInfoByte))

	if error != nil {
		uc.HandleErrorResponse(c, errors.New("生成 token 失败"))
		return
	}

	uc.HandleSuccessResponse(c, tokenString)
}

// 获取用户信息
func (uc UserController) GetUserInfo(c *gin.Context) {
	//userId, error := c.Param("userId")
}