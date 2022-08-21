package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	if err := mysql.CheckUserExist(p.UserName); err != nil {
		return err
	}

	// 2.生成UID
	userID := snowflake.GetID()
	// 3.保存进数据库
	user := &models.User{
		UserID:   userID,
		Username: p.UserName,
		Password: p.Password,
	}

	return mysql.InsertUser(user)

}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)

}
