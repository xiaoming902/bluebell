package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// FollowHandler 关注/取关功能 1 = 关注 0 = 取关

func FollowHandler(c *gin.Context) {

	p := new(models.ParamFollow)
	p.Userid, _ = getCurrentUser(c)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("FollowHandler invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	err := logic.Follow(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "添加成功")

}

func GetFollowersHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	data, err := logic.GetFollowers(userId)

	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)

}

func GetFollowingHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	followData, err := logic.GetFollowing(userId)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, followData)

}
