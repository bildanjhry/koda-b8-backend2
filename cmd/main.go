package main

import (
	"github.com/bildanjhry/auth/internal/di"
	"github.com/bildanjhry/go_shared-lib/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	c := di.NewContainer()
	userHandler := c.UserHandler()
	utils.LoadEnv()

	r.POST("/register", userHandler.Create)
	r.POST("/login", userHandler.Login)

	r.Run("0.0.0.0:8080")
}
