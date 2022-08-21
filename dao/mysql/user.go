package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "jichen"

func CheckUserExist(username string) (err error) {
	sqlstr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlstr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlstr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))

}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlstr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlstr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}

	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return

}

func GetUserByID(idStr string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, idStr)
	return
}
