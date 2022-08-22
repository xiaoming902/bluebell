package models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UseInfo struct {
	Username    string `db:"username"`
	Description string `db:"description"`
}

type Userid struct {
	UserID       int64 `db:"user_id" json:"user_id"`
	FollowUserId int64 `db:"follow_user_id" json:"follow_user_id"`
}
