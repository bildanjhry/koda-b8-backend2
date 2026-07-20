package repo

import "github.com/bildanjhry/auth/internal/model"

type UserRepo struct {
	dataUser *[]model.Users
}

func NewUserRepo(dataUser *[]model.Users) *UserRepo {
	return &UserRepo{
		dataUser: dataUser,
	}
}

func (u *UserRepo) Create(data *model.UserForm) *model.Users {
	id := int64(len(*u.dataUser) + 1)
	newUser := model.Users{
		Id:       id,
		Email:    data.Email,
		Password: data.Password,
	}
	*u.dataUser = append(*u.dataUser, newUser)
	return &newUser
}
