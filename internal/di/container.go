package di

import (
	"github.com/bildanjhry/auth/internal/handler"
	"github.com/bildanjhry/auth/internal/model"
	"github.com/bildanjhry/auth/internal/repo"
	"github.com/bildanjhry/auth/internal/service"
)

type Container struct {
	userData    *[]model.Users
	userRepo    *repo.UserRepo
	userService *service.UserService
	userHandler *handler.UserHandler
}

func (c *Container) initDeps() {
	c.userRepo = repo.NewUserRepo(c.userData)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func NewContainer() *Container {
	userData := []model.Users{}

	container := &Container{
		userData: &userData,
	}
	container.initDeps()
	return container
}
