package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func PostVoteController(c *gin.Context) {
	err := c.ShouldBindQuery("name")
	if err != nil {
		return
	}
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			//ResponseError(c, CodeInvalidParam)
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		fmt.Println(err)
		return
	}
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost", zap.Error(err))
	}
	ResponseSuccess(c, nil)

}
