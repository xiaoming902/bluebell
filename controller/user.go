package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

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
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, token)

}

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,: ", err.Error())
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}

//GetUserInfoHandler 获取用户信息
func GetUserInfoHandler(c *gin.Context) {

	userid := c.Param("userid")

	date, err := logic.GetUserInfo(userid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
		return
	}

	ResponseSuccess(c, date)

}

func ChangeUserPassword(c *gin.Context) {

	p := new(models.ChangeUserPassword)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//密码检查
	err := logic.CheckPassword(p.Password)
	if err != nil {
		ResponseError(c, PasswordLengthLimit)
		return
	}

	//旧密码校验

	u, _ := c.Get("username")

	if !logic.ValidPassword(u.(string), p.OldPassword) {
		ResponseError(c, CodeOldPassword)
		return

	}

	ResponseSuccess(c, u)

}
