package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.CommunityHandler() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

func CommunityDetailHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, errs := strconv.ParseInt(idStr, 10, 64)
	if errs != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.CommunityDetailHandler() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)

}
