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
