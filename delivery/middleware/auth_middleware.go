package middleware

import (
	"fmt"
	"strings"

	"github.com/St0rage/Simpan-Uang/service"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/St0rage/Simpan-Uang/utils/authenticator"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken() gin.HandlerFunc
	PiggyBankAuthorization() gin.HandlerFunc
	WishlistAuthorization() gin.HandlerFunc
}

type authMiddleware struct {
	tokenServ        authenticator.AccessToken
	piggyBankService service.PiggyBankService
	wishlistService service.WishlistService
}

func NewAuthMiddleware(tokenServ authenticator.AccessToken, piggyBankService service.PiggyBankService, wishlistService service.WishlistService) AuthMiddleware {
	return &authMiddleware{
		tokenServ:        tokenServ,
		piggyBankService: piggyBankService,
		wishlistService: wishlistService,
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

func (auth *authMiddleware) PiggyBankAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		piggyBankId := ctx.Param("piggyBankId")
		userId := fmt.Sprintf("%v", ctx.MustGet("id"))

		piggyBankUserId, err := auth.piggyBankService.GetPiggyBankUser(piggyBankId)
		if err != nil {
			ctx.AbortWithStatusJSON(404, utils.ResponseWrapper{
				Code:   404,
				Status: "NOT FOUND",
				Data:   nil,
			})
			return
		}

		if piggyBankUserId == userId {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(404, utils.ResponseWrapper{
				Code:   404,
				Status: "NOT FOUND",
				Data:   nil,
			})
			return
		}

	}
}

func (auth *authMiddleware) WishlistAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wishlistId := ctx.Param("wishlistId")
		userId := fmt.Sprintf("%v", ctx.MustGet("id"))

		wishlistUserId, err := auth.wishlistService.GetWishlistUser(wishlistId)
		if err != nil {
			ctx.AbortWithStatusJSON(404, utils.ResponseWrapper{
				Code:   404,
				Status: "NOT FOUND",
				Data:   nil,
			})
			return
		}

		if wishlistUserId == userId {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(404, utils.ResponseWrapper{
				Code:   404,
				Status: "NOT FOUND",
				Data:   nil,
			})
			return
		}

	}
}