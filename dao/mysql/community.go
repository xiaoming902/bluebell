package mysql

import (
	"bluebell/models"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	fmt.Println(err)
	return
}

func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {

	community := new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	//err := db.Get(community, sqlStr, id)

	//if err := db.Get(community, sqlStr, id); err != nil {
	//	err = ErrorInvalidID
	//}
	err := db.Get(community, sqlStr, id)
	return community, err

}

func GetCommunityByID(idStr string) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorInvalidID
		return
	}
	return
}
