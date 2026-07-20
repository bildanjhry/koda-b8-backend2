package model

type Users struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
