package di

import (
	"github.com/bildanjhry/auth/internal/handler"
	"github.com/bildanjhry/auth/internal/repo"
	"github.com/bildanjhry/auth/internal/service"
)

type Container struct {
	userRepo    *repo.UserRepo
	userService *service.UserService
	userHandler *handler.UserHandler
}

func (c *Container) initDeps() {
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func NewContainer() *Container {

	container := &Container{}
	container.initDeps()
	return container
}
