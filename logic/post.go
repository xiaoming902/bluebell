package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	p.ID = snowflake.GetID()

	return mysql.CreatePost(p)
}

func GetPost(postID string) (post *models.ApiPostDetail, err error) {
	post, err = mysql.GetPostByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed", zap.String("post_id", postID), zap.Error(err))
		return nil, err
	}
	user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		return
	}
	post.AuthorName = user.Username
	community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		return
	}
	post.AuthorName = community.Name
	return post, nil
}

func GetPostList() (data []*models.ApiPostDetail, err error) {

	postList, err := mysql.GetPostList()
	if err != nil {
		fmt.Println(err)
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
			continue
		}
		post.AuthorName = user.Username
		community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
			continue
		}
		post.AuthorName = community.Name
		data = append(data, post)
	}
	return
}
