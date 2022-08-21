package models

type Follow struct {
	Id           int64 `db:"id"`
	UserId       int64 `db:"user_id"`
	FollowUserId int64 `db:"follow_user_id"`
	IsValid      int8  `db:"is_valid"`
}
