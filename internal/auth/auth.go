package auth

import (
	"errors"
	"medium-be/internal/constants"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractToken(e echo.Context) (Claims, error) {
	user := e.Get("user").(*jwt.Token)
	claims := &Claims{}
	jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_ACCESS_KEY), nil
	})

	return *claims, errors.New("invalid token")
}
