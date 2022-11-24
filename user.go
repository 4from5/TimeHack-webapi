package webApi

type User struct {
	Id       int    `json:"-" db:"user_id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}
