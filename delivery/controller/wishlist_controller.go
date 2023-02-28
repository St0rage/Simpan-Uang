package controller

import (
	"fmt"

	"github.com/St0rage/Simpan-Uang/delivery/middleware"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/service"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/gin-gonic/gin"
)

type WishlistController struct {
	router  *gin.Engine
	service service.WishlistService
}

func (wc *WishlistController) GetWishlist(ctx *gin.Context) {
	Userid := fmt.Sprintf("%v", ctx.MustGet("id"))

	user, err := wc.service.GetWishlist(Userid)
	if err != nil {
		utils.HandleInternalServerError(ctx)
	} else {
		utils.HandleSuccess(ctx, user)
	}
}

func (wc *WishlistController) CreateNewWishlist(ctx *gin.Context) {
	Userid := fmt.Sprintf("%v", ctx.MustGet("id"))
	var wishlist *web.WishlistRequest
	err := ctx.ShouldBindJSON(&wishlist)
	if err != nil {
		utils.HandleInternalServerError(ctx)
	} else {
		wc.service.CreateNewWishlist(Userid,wishlist)
		utils.HandleSuccessCreated(ctx, wishlist)
	}
}

// func (uc *UserController) registerUser(ctx *gin.Context) {
// 	var user *domain.User
// 	err := ctx.ShouldBindJSON(&user)
// 	if err != nil {
// 		utils.HandleBadRequest(ctx, err.Error())
// 	} else {
// 		err := uc.service.Register(user)
// 		if err != nil {
// 			utils.HandleInternalServerError(ctx)
// 		} else {
// 			utils.HandleSuccessCreated(ctx, map[string]string{
// 				"message": "User berhasil dibuat",
// 			})
// 		}
// 	}
// }

// func (uc *UserController) loginUser(ctx *gin.Context) {
// 	var loginRequest *web.UserLoginRequest
// 	err := ctx.ShouldBindJSON(&loginRequest)
// 	if err != nil {
// 		utils.HandleBadRequest(ctx, err.Error())
// 	} else {
// 		token, err := uc.service.Login(loginRequest)
// 		if err != nil {
// 			utils.HandleBadRequest(ctx, map[string]string{
// 				"message": "Email atau Password salah",
// 			})
// 		} else {
// 			utils.HandleSuccess(ctx, map[string]string{
// 				"token": token,
// 			})
// 		}
// 	}
// }

// func (uc *UserController) forgotPassword(ctx *gin.Context) {
// 	var resetRequest *web.UserResetRequest
// 	err := ctx.ShouldBindJSON(&resetRequest)
// 	if err != nil {
// 		utils.HandleBadRequest(ctx, err.Error())
// 	} else {
// 		err := uc.service.ForgotPassword(resetRequest)
// 		if err != nil {
// 			utils.HandleNotFound(ctx, map[string]string{
// 				"message": "Email tidak ditemukan",
// 			})
// 		} else {
// 			utils.HandleSuccess(ctx, map[string]string{
// 				"message": "Berhasil reset password, cek email anda",
// 			})
// 		}
// 	}
// }

func NewWishlistController(r *gin.Engine, service service.WishlistService, authMdw middleware.AuthMiddleware) *WishlistController {
	controller := WishlistController{
		router:  r,
		service: service,
	}

	// userRouteGroup := controller.router.Group("/api/user")
	// userRouteGroup.GET("/", authMdw.RequireToken(), controller.getUser)
	// userRouteGroup.POST("/register", controller.registerUser)
	// userRouteGroup.POST("/login", controller.loginUser)
	// userRouteGroup.POST("/forgot-password", controller.forgotPassword)

	wishlistRouteGroup := controller.router.Group("/api/wishlist", authMdw.RequireToken())
	wishlistRouteGroup.GET("/", controller.GetWishlist)
	wishlistRouteGroup.POST("/create", controller.CreateNewWishlist)
	return &controller
}
