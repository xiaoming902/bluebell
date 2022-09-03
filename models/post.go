package models

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name" db:"username"`
	*Post
	*CommunityDetail `json:"community_detail"`
}

type PostID struct {
	ID       int64 `json:"id" db:"post_id"`
	AuthorId int64 `db:"author_id"`
}
