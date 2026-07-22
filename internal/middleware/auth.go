package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bildanjhry/auth/internal/lib"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		fmt.Println(authHeader)
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &lib.Response{
				Success: false,
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			return
		}

		token, _ := strings.CutPrefix(authHeader, "Bearer ")
		if lib.VerifyToken(token) {
			ctx.Next()
			return
		}

		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
