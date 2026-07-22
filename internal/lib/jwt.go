package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTPayload struct {
	IdUser int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int64) string {
	claims := JWTPayload{
		id,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		fmt.Println(err.Error())
	}
	return ss
}

func VerifyToken(token string) bool {
	_, err := jwt.ParseWithClaims(token, &JWTPayload{},
		func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
