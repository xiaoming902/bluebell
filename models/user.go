package models

type User struct {
	UserID      int64  `db:"user_id"`
	Username    string `db:"username"`
	Phone       string `json:"phone" db:"phone"`
	Password    string `db:"password"`
	Salt        string `json:"salt"`
	Description string `db:"description"`
	Status      int    `json:"status"`
	IsAdmin     bool   `json:"is_admin"`
}

type UseInfo struct {
	Username    string `db:"username"`
	Description string `db:"description"`
}

type Userid struct {
	Username string `db:"username" json:"user_name"`
	UserID   int64  `db:"follow_user_id" json:"user_id"`
}

type UserId struct {
	Username string `db:"username" json:"user_name"`
	UserID   int64  `db:"user_id" json:"user_id"`
}
