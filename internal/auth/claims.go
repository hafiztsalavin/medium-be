package auth

import (
	"github.com/golang-jwt/jwt"
)

// result as an auth token
type Claims struct {
	Id    uint32
	Email string
	jwt.StandardClaims
}
