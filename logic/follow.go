package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"errors"
)

func Follow(p *models.ParamFollow) error {

	t := p.Act

	err, _ := mysql.CheckFollow(p)

	if errors.Is(err, mysql.ErrorInvalid) {
		return err
	}

	if p.Userid == p.Fid {
		return errors.New("不能关注自己")
	}

	if err == nil && *p.Act == 1 {
		err = mysql.SaveFollow(p)
		if err != nil {
			return err
		}
	}

	// 如果有关注信息，且目前处于关注状态，且要进行取关操作 -- 取关操作
	if errors.Is(err, mysql.ErrorFollowExist) && *p.Act == 1 && *t == 0 {
		err = mysql.UpdateFollow(p, t)
		if err != nil {
			return err
		}
	}

	// 如果有关注信息，且目前处于取关状态，且要进行关注操作 -- 重新关注
	if errors.Is(err, mysql.ErrorFollowExist) && *p.Act == 0 && *t == 0 {
		err = mysql.UpdateFollow(p, t)
		if err != nil {
			return err
		}
	}

	return err

}

// GetFollowers 查看关注我的人 (粉丝)
func GetFollowers(userId int64) (*models.Userid, error) {

	return mysql.GetFollowers(userId)

}

func GetFollowing(userId int64) (*map[string]interface{}, error) {

	followData, err := mysql.GetFollowing(userId)
	data := make(map[string]interface{}, 0)
	data["count"] = len(followData)
	data["data"] = followData

	return &data, err

}
