package authenticator

import "github.com/golang-jwt/jwt/v5"

type MyClaims struct {
	jwt.RegisteredClaims
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
