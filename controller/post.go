package controller

import (
	"bluebell/dao/redis"
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const (
	orderTime  = "time"
	orderScore = "score"
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
		ResponseError(c, CodeServerBusy)
		return
	}
	err = redis.CreatePost(p.ID)

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

	//data, err := logic.GetPost(strconv.FormatInt(pid, 10))
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)

}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)

	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)

}

//func GetPostListHandler2(c *gin.Context) {
//	// Get请求参数(query string): /api/v1/posts2?page=1&size=10&order=time
//	p := &models.ParamPostList{
//		Page:  1,
//		Size:  10,
//		Order: orderTime,
//	}
//	if err := c.ShouldBindQuery(p); err != nil {
//		ResponseError(c, CodeInvalidParam)
//	}
//}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
