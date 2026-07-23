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
	return res, nil
}

func (r *UserService) GetAll(par *model.UserParams) ([]*model.Users, error) {
	res := r.repo.GetAll(par)
	return res, nil
}

func (r *UserService) GetById(id *int64) (*model.Users, error) {
	res, err := r.repo.GetById(id)
	return res, err
}

func (r *UserService) Delete(id *int64) error {
	res := r.repo.Delete(id)
	return res
}

func (r *UserService) Edit(id *int64, data *model.UserEmail) (*model.Users, error) {
	res, err := r.repo.Edit(id, data)
	return res, err
}

func (r *UserService) UploadPicture(id *int64, data *model.UserPicture) (*model.Users, error) {
	res, err := r.repo.UploadPicture(id, data)
	return res, err
}

func (r *UserService) Login(data *model.LoginForm) (*model.Users, error) {
	if len(data.Password) < 5 {
		return &model.Users{}, errors.New("Minimum password length 5")
	}
	res, err := r.repo.Login(data)
	if err != nil {
		return &model.Users{}, err
	}
	return &model.Users{
		Id:        res.Id,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
	}, nil
}
