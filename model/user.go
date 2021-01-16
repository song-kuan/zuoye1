package model

type User struct {
    Id       int64  `xorm:"pk autoincr" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}