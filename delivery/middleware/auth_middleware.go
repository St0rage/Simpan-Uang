package middleware

import (
	"strings"

	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/St0rage/Simpan-Uang/utils/authenticator"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authMiddleware struct {
	tokenServ authenticator.AccessToken
}

func NewAuthMiddleware(tokenServ authenticator.AccessToken) AuthMiddleware {
	return &authMiddleware{
		tokenServ: tokenServ,
	}
}

func (auth *authMiddleware) RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.Request.Header.Get("Authorization")

		tokenString := strings.Replace(authToken, "Bearer ", "", -1)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, utils.ResponseWrapper{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}

		token, err := auth.tokenServ.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(401, utils.ResponseWrapper{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}

		if token != nil {
			ctx.Set("id", token["id"])
			ctx.Set("name", token["name"])
			ctx.Set("email", token["email"])
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(401, utils.ResponseWrapper{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}
	}
}
