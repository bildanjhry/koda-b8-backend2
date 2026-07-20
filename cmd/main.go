package main

import (
	"github.com/bildanjhry/auth/internal/di"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	c := di.NewContainer()
	userHandler := c.UserHandler()
	r.POST("/reg", userHandler.Create)

	r.Run("0.0.0.0:8080")
}
