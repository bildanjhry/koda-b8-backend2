package main

import (
	"github.com/bildanjhry/auth/internal/di"
	"github.com/bildanjhry/auth/internal/middleware"
	"github.com/bildanjhry/go_shared-lib/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Backend Auth
//	@version		1.0
//	@description	This is a sample server celler server.

func main() {
	r := gin.Default()
	c := di.NewContainer()
	userHandler := c.UserHandler()
	utils.LoadEnv()
	r.Static("/uploads", "./uploads")

	auth := r.Group("/auth")
	auth.Use(middleware.Cors())
	auth.OPTIONS("/login", func(ctx *gin.Context) {})
	auth.POST("/register", userHandler.Create)
	auth.OPTIONS("/register", func(ctx *gin.Context) {})
	auth.POST("/login", userHandler.Login)

	users := r.Group("/user")
	users.Use(middleware.Cors())
	users.Use(middleware.Auth())
	users.OPTIONS("/all", func(ctx *gin.Context) {})
	users.GET("/all", userHandler.GetAll)
	users.OPTIONS("/delete/:id", func(ctx *gin.Context) {})
	users.DELETE("/delete/:id", userHandler.Delete)
	users.OPTIONS("/edit/:id", func(ctx *gin.Context) {})
	users.PATCH("/edit/:id", userHandler.GetById)
	users.OPTIONS("/detail/:id", func(ctx *gin.Context) {})
	users.GET("/detail/:id", userHandler.GetById)
	users.OPTIONS("/upload-pic/:id", func(ctx *gin.Context) {})
	users.PATCH("/upload-pic/:id", userHandler.UploadPicture)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("0.0.0.0:8080")
}
