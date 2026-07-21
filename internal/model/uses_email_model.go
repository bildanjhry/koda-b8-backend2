package model

type UserEmail struct {
	Email string `json:"email" form:"email" binding:"email,required"`
}
