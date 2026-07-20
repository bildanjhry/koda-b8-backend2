package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/bildanjhry/auth/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	dataUser *[]model.Users
}

func NewUserRepo(dataUser *[]model.Users) *UserRepo {
	return &UserRepo{
		dataUser: dataUser,
	}
}

func (u *UserRepo) Create(data *model.UserForm) *model.Users {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	response, err := pool.Query(context.Background(), `INSERT INTO "users" ("email", "password") VALUES ($1, $2) RETURNING "id", "email", "password", "created_at", "updated_at"`, data.Email, data.Password)

	users, formErr := pgx.CollectOneRow(response, pgx.RowToAddrOfStructByName[model.Users])

	if formErr != nil {
		return &model.Users{}
	}

	newUser := model.Users{
		Id:        users.Id,
		Email:     users.Email,
		Password:  users.Password,
		CreatedAt: users.CreatedAt,
	}

	return &newUser
}
