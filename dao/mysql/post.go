package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {

	sqlStr := `insert into post (post_id, title, content, author_id, community_id) value (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	return
}

func GetPostById(pid int64) (post *models.ApiPostDetail, err error) {

	post = new(models.ApiPostDetail)

	//sqlStr := `select post_id, title, content, author_id, community_id, create_time FROM post where post_id = ?`

	sqlStr := `select post_id,title, content, author_id, user.username ,post.community_id, post.create_time ,c.community_name, c.introduction FROM post LEFT JOIN user on post.author_id = user.user_id LEFT JOIN community c on post.community_id = c.community_id where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return

}

func GetPostList() (posts []*models.ApiPostDetail, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	limit 2
	`
	posts = make([]*models.ApiPostDetail, 0, 2)
	err = db.Select(&posts, sqlStr)
	return

}

func GetPostByID(idStr string) (post *models.ApiPostDetail, err error) {
	post = new(models.ApiPostDetail)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`
	err = db.Get(post, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorInvalidID
		return
	}
	return
}
