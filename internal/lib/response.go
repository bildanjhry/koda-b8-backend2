package lib

import "time"

type Response struct {
	Success bool
	Status  int
	Message string
	Results any
}

type LoginResponse struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt *time.Time
	Token     string
}
