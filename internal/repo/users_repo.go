package repo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/bildanjhry/auth/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	dataUser []*model.Users
}

func NewUserRepo(dataUser []*model.Users) *UserRepo {
	return &UserRepo{
		dataUser: dataUser,
	}
}

func (u *UserRepo) Login(data *model.LoginForm) (*model.Users, error) {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	response, resErr := pool.Query(context.Background(),
		`SELECT "id", "name", "email", "password", "created_at", "updated_at", "picture" FROM "users" 
		WHERE email=$1 AND password=$2`,
		data.Email, data.Password)

	if resErr != nil {
		return &model.Users{}, resErr
	}

	users, formErr := pgx.CollectOneRow(response, pgx.RowToAddrOfStructByName[model.Users])

	if formErr != nil {
		return nil, formErr
	}

	newUser := model.Users{
		Id:        users.Id,
		Name:      users.Name,
		Email:     users.Email,
		Password:  users.Password,
		CreatedAt: users.CreatedAt,
		Picture:   users.Picture,
	}

	return &newUser, nil
}

func (u *UserRepo) GetById(id *int64) (*model.Users, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	response, resErr := pool.Query(context.Background(),
		`SELECT "id", "name", "email", "password", "created_at", "updated_at", "picture" FROM "users" WHERE id=$1`,
		id)
	if resErr != nil {
		fmt.Println(resErr.Error())
		return nil, resErr
	}

	users, formatErr := pgx.CollectOneRow(response, pgx.RowToAddrOfStructByName[model.Users])
	if formatErr != nil {
		return nil, formatErr
	}

	return users, nil
}

func (u *UserRepo) GetAll(par *model.UserParams) []*model.Users {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	intPage, _ := strconv.ParseInt(par.PAGE, 10, 32)
	intLimit, _ := strconv.ParseInt(par.LIMIT, 10, 32)
	var page int64 = 1
	page = (intPage * intLimit) - intLimit

	qEmail := "%" + par.EMAIL + "%"
	qName := "%" + par.NAME + "%"
	querySearch := ""

	query := `SELECT "id", "name", "email", "password", "created_at", "updated_at", "picture" 
	FROM "users" `

	if par.NAME != "" && par.EMAIL != "" {
		page -= (page)
		querySearch = fmt.Sprintf(`WHERE name ILIKE '%s' OR email ILIKE '%s'`, qName, qEmail)
	} else if par.NAME != "" {
		page -= (page)
		querySearch = fmt.Sprintf(`WHERE name ILIKE '%s'`, qName)
	} else if par.EMAIL != "" {
		page -= (page)
		querySearch = fmt.Sprintf(`WHERE email ILIKE '%s'`, qEmail)
	}

	pagination := fmt.Sprintf(` ORDER BY %s %s LIMIT %s OFFSET %d`, par.ORDER_BY, par.ORDER, par.LIMIT, page)

	fmt.Println(query + querySearch + pagination)

	response, errRes := pool.Query(context.Background(), query+querySearch+pagination)
	if errRes != nil {
		fmt.Println(errRes.Error())
	}
	users, formErr := pgx.CollectRows(response, pgx.RowToAddrOfStructByName[model.Users])
	if formErr != nil {
		fmt.Println(formErr.Error())
	}
	return users
}

func (u *UserRepo) Delete(id *int64) error {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	response, resErr := pool.Exec(context.Background(), `DELETE FROM "users" WHERE id=$1`, id)
	if response.RowsAffected() != 1 {
		return resErr
	}
	return nil
}

func (u *UserRepo) Edit(id *int64, data *model.UserEmail) (*model.Users, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	response, resErr := pool.Query(context.Background(),
		`UPDATE "users" SET email=$1, updated_at=NOW() WHERE id=$2 
		RETURNING "id", "email", "password", "created_at", "updated_at"`, data.Email, id)
	if resErr != nil {
		return &model.Users{}, resErr
	}
	users, formErr := pgx.CollectOneRow(response, pgx.RowToAddrOfStructByName[model.Users])
	if formErr != nil {
		return &model.Users{}, resErr
	}
	return users, nil

}

func (u *UserRepo) UploadPicture(id *int64, data *model.UserPicture) (*model.Users, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	response, resErr := pool.Query(context.Background(),
		`UPDATE "users" SET picture=$1, updated_at=NOW() WHERE id=$2 
		RETURNING "id", "email", "picture", "password", "created_at", "updated_at"`, data.Picture, id)
	if resErr != nil {
		return &model.Users{}, resErr
	}
	users, formErr := pgx.CollectOneRow(response, pgx.RowToAddrOfStructByName[model.Users])
	if formErr != nil {
		return &model.Users{}, resErr
	}
	return users, nil

}

func (u *UserRepo) Create(data *model.UserForm) *model.Users {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success connected")
	}

	defer pool.Close()

	response, err := pool.Query(context.Background(),
		`INSERT INTO "users" ("name", "email", "password") 
		VALUES ($1, $2, $3) RETURNING "id", "name", "email", "password", "created_at", "updated_at"`,
		data.Name, data.Email, data.Password)

	users, formErr := pgx.CollectOneRow(response, pgx.RowToAddrOfStructByName[model.Users])

	if formErr != nil {
		return &model.Users{}
	}

	fmt.Println(users)

	return users
}

func (u *UserRepo) GetByAttrs(data *string) ([]*model.Users, error) {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer pool.Close()

	input := fmt.Sprintf("%s", *data)
	input = input + string("%")

	response, resErr := pool.Query(context.Background(),
		`SELECT * FROM "users" WHERE name LIKE $1 OR email LIKE $1 `,
		input)
	if resErr != nil {
		fmt.Println(resErr.Error())
		return nil, resErr
	}

	users, formErr := pgx.CollectRows(response, pgx.RowToAddrOfStructByName[model.Users])
	if formErr != nil {
		fmt.Println(formErr.Error())
		return nil, formErr
	} else if len(users) < 1 {
		return nil, errors.New("User not found")
	}

	return users, nil
}
