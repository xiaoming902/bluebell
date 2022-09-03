package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/errcode"
	"bluebell/pkg/snowflake"
	"fmt"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	p.ID = snowflake.GetID()

	return mysql.CreatePost(p)
}

//func GetPost(postID string) (post *models.ApiPostDetail, err error) {
//	post, err = mysql.GetPostByID(postID)
//	if err != nil {
//		zap.L().Error("mysql.GetPostByID(postID) failed", zap.String("post_id", postID), zap.Error(err))
//		return nil, err
//	}
//	user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
//	if err != nil {
//		zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
//		return
//	}
//	post.AuthorName = user.Username
//	community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
//	if err != nil {
//		zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
//		return
//	}
//	post.AuthorName = community.Name
//	return post, nil
//}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {

	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		fmt.Println(err)
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
			continue
		}
		post.AuthorName = user.Username
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
			continue
		}
		post.CommunityDetail = community
		data = append(data, post)
	}
	return
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {

	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}

	// 根据作者id查询作者信息
	user, err := mysql.GetUserByID(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorId) failed",
			zap.Int64("author_id", post.AuthorId),
			zap.Error(err))
		return
	}

	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
			zap.Int64("CommunityID", post.CommunityID),
			zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return

}

func UpdatePost(p *models.Post, pid int64) error {
	id, err := GetPostById(pid)
	if err != nil {
		return errcode.NoPermission
	}
	// 判断当前id是否为帖子创建者id
	if id.AuthorId != p.AuthorId {
		return errcode.NoPermission
	}
	p.ID = pid
	err = mysql.UpdatePost(p)
	if err != nil {
		return errcode.ServerError
	}

	return nil
}

func DeletePost(p *models.PostID) error {
	id, err := GetPostById(p.ID)
	if err != nil {
		return errcode.NoPermission
	}
	// 判断当前id是否为帖子创建者id
	if id.AuthorId != p.AuthorId {
		return errcode.NoPermission
	}

	err = mysql.DeletePost(p)
	if err != nil {
		return errcode.ServerError
	}

	return nil
}
