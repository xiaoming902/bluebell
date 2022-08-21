package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)

}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	token, err := logic.Login(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	ResponseSuccess(c, token)

}
