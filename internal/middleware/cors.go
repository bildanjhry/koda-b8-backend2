package middleware

import (
	"net/http"

	"github.com/bildanjhry/auth/internal/lib"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}

		// 		if ctx.Request.Method == "OPTIONS" {
		// 			ctx.Status(http.StatusOK)
		// 		}

		if ctx.GetHeader("Authorization") != "hello" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &lib.Response{
				Success: false,
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			return
		}

		ctx.Next()
	}
}
