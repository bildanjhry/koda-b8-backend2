package model

import "time"

type Users struct {
	Id        int64     `json:"id"`
	Name      *string   `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	Picture   *string   `json:"picture"`
}
