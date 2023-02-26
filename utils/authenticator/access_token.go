package authenticator

import (
	"fmt"
	"log"

	"github.com/St0rage/Simpan-Uang/config"
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/golang-jwt/jwt/v5"
)

type AccessToken interface {
	CreateAccessToken(cred *domain.User) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type accessToken struct {
	cfg config.TokenConfig
}

func NewAccessToken(config config.TokenConfig) AccessToken {
	return &accessToken{
		cfg: config,
	}
}

func (accessToken *accessToken) CreateAccessToken(cred *domain.User) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: accessToken.cfg.ApplicationName,
		},
		Id:    cred.Id,
		Name:  cred.Name,
		Email: cred.Email,
	}
	token := jwt.NewWithClaims(accessToken.cfg.JwtSigningMethod, claims)

	return token.SignedString([]byte(accessToken.cfg.JwtSignatureKey))
}

func (accessToken *accessToken) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != accessToken.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(accessToken.cfg.JwtSignatureKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != accessToken.cfg.ApplicationName {
		log.Println("Token Invalid")
		return nil, err
	}
	return claims, nil
}
