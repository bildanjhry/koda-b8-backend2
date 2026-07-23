package model

type LoginForm struct {
	Email    string `json:"email" form:"email" binding:"email,required"`
	Password string `json:"password" form:"password" binding:"required"`
}
