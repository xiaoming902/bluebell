package models

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction,string" binding:"required, oneof=1 0 -1"`
}

type ParamFollow struct {
	Userid int64 `db:"user_id"`
	Fid    int64 `json:"fid" db:"follow_user_id" binding:"required"`
	Act    *int  `json:"act" db:"is_valid" binding:"required,oneof=0 1"`
}
