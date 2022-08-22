package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 用户注册
	r.POST("/signup", controller.SignUpHandler)

	// 用户登录
	r.POST("/login", controller.LoginHandler)

	// 获取验证码
	r.GET("/captcha", controller.GetCaptcha)

	// 查看用户主页
	r.GET("/user/:userid", controller.GetUserInfoHandler)

	// 查看关注我的人 (粉丝)

	r.GET("/follow/followers", controller.GetFollowersHandler)

	// 查看我关注的人

	r.GET("/follow/following", controller.GetFollowingHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/post", controller.GetPostListHandler)

		//帖子投票
		v1.POST("/vote", controller.PostVoteController)

		// 关注用户
		v1.POST("/follow", controller.FollowHandler)

		// 修改密码
		v1.POST("/user/password", controller.ChangeUserPassword)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})

	})
	return r
}
