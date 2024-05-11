package handler

import (
	"log"
	"time"

	"github.com/Erwin011895/TaskManagementApp/config"
	"github.com/Erwin011895/TaskManagementApp/model"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(conf *config.Config, user model.User) (string, error) {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTCustomClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	})

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(conf.App.Key))
	return t, err
}

func GetClaim(c *gin.Context) *JWTCustomClaims {
	token := c.GetHeader("Authorization")
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return "", nil
	})
	if err != nil {
		log.Println(err)
	}
	return t.Claims.(*JWTCustomClaims)
}
