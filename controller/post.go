package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {

	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUser(c)

	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	p.AuthorId = userID

	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("sd")
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)

}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPost(strconv.FormatInt(pid, 10))
	if err != nil {
		zap.L().Error("logic.GetPostByID ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)

}

func GetPostListHandler(c *gin.Context) {

	_, err := logic.GetPostList()

	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

}
