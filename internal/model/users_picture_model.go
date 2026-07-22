package model

type UserPicture struct {
	Picture string `json:"picture" form:"picture" binding:"required"`
}
