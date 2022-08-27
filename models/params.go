package models

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required, eqfield=Password"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	//UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction" binding:"oneof=1 0 -1"`
}

type ParamFollow struct {
	Userid int64 `db:"user_id"`
	Fid    int64 `json:"fid" db:"follow_user_id" binding:"required"`
	Act    *int  `json:"act" db:"is_valid" binding:"required,oneof=0 1"`
}

type ChangeUserPassword struct {
	Password    string `form:"password" json:"password" binding:"required"`
	OldPassword string `form:"old_password" json:"old_password" binding:"required"`
}
