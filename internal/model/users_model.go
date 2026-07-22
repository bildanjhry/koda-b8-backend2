package model

import "time"

type Users struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Picture   *string   `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
