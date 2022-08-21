package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

func CheckFollow(data *models.ParamFollow) (error, int) {

	var Userid struct {
		UserID int64 `db:"user_id"`
	}

	sqlstr := `select user_id from user where user_id = ?`

	if err := db.Get(&Userid, sqlstr, data.Fid); err != nil {
		if err == sql.ErrNoRows {
			return ErrorInvalid, 2
		} else {
			zap.L().Error("查询失败 CheckFollow")
			return ErrorInvalid, 2
		}
	}

	sqlstr = `select user_id, follow_user_id, is_valid  from follow where user_id = ? and follow_user_id = ?`

	err := db.Get(data, sqlstr, data.Userid, data.Fid)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 2
		} else {
			zap.L().Error("查询失败 CheckFollow")
			return ErrorInvalid, 2
		}
	}

	return ErrorFollowExist, *data.Act

}

func SaveFollow(data *models.ParamFollow) error {

	sqlstr := `insert into follow(user_id, follow_user_id, is_valid) values(?,?,?)`
	_, err := db.Exec(sqlstr, data.Userid, data.Fid, data.Act)
	return err

}

func UpdateFollow(data *models.ParamFollow, t *int) error {

	sqlstr := `update follow set is_valid = ? where follow_user_id = ? and user_id = ?`
	_, err := db.Exec(sqlstr, t, data.Fid, data.Userid)

	return err

}
