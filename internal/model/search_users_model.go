package model

type Search struct {
	Input string `json:"search" form:"search" binding:"required"`
}
