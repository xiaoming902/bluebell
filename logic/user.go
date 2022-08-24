package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
	"github.com/gofrs/uuid"
	"unicode/utf8"
)

// EncryptPasswordAndSalt 密码加密&生成salt
func EncryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = mysql.EncryptPassword(mysql.EncryptPassword(password) + salt)

	return password, salt
}

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	if err := mysql.CheckUserExist(p.UserName); err != nil {
		return err
	}

	// 2.生成UID && 密码加密&生成salt
	userID := snowflake.GetID()

	password, salt := EncryptPasswordAndSalt(p.Password)

	// 3.保存进数据库
	user := &models.User{
		UserID:   userID,
		Username: p.UserName,
		Password: password,
		Salt:     salt,
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

// CheckPassword 密码检查
func CheckPassword(password string) error {
	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 16 {
		return errors.New("密码长度6 ~ 16")
	}
	return nil
}

func ValidPassword(user, password string) bool {

	if mysql.EncryptPassword(password) == mysql.CheckPassword(user) {
		return true
	}

	return false

}

func GetUserInfo(userid string) (*models.UseInfo, error) {

	return mysql.GetUserInfo(userid)

}
