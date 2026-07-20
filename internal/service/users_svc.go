package service

import (
	"errors"

	"github.com/bildanjhry/auth/internal/model"
	"github.com/bildanjhry/auth/internal/repo"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (r *UserService) Create(data *model.UserForm) (*model.Users, error) {
	if len(data.Password) < 5 {
		return &model.Users{}, errors.New("Minimum password length 5")
	}
	res := r.repo.Create(data)
	return &model.Users{
		Id:    res.Id,
		Email: res.Email,
	}, nil
}
