package controller

import (
	"fmt"

	"github.com/St0rage/Simpan-Uang/delivery/middleware"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/service"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	router  *gin.Engine
	service service.UserService
}

func (uc *UserController) getUser(ctx *gin.Context) {
	id := fmt.Sprintf("%v", ctx.MustGet("id"))

	user := uc.service.GetUser(id)

	utils.HandleSuccess(ctx, user)
}

func (uc *UserController) registerUser(ctx *gin.Context) {
	var newUserRequest *web.UserRegisterRequest
	err := ctx.ShouldBindJSON(&newUserRequest)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := uc.service.Register(newUserRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, map[string]string{
				"message": err.Error(),
			})
		} else {
			utils.HandleSuccessCreated(ctx, map[string]string{
				"message": "User berhasil dibuat",
			})
		}
	}
}

func (uc *UserController) loginUser(ctx *gin.Context) {
	var loginRequest *web.UserLoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		token, err := uc.service.Login(loginRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, map[string]string{
				"message": "Email atau Password salah",
			})
		} else {
			utils.HandleSuccess(ctx, map[string]string{
				"message": "Login Berhasil",
				"token":   token,
			})
		}
	}
}

func (uc *UserController) forgotPassword(ctx *gin.Context) {
	id := fmt.Sprintf("%v", ctx.MustGet("id"))

	isAdmin := uc.service.CheckAdmin(id)

	if !isAdmin {
		utils.HandleUnauthorized(ctx)
	} else {
		var resetRequest *web.UserResetRequest
		err := ctx.ShouldBindJSON(&resetRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, err.Error())
		} else {
			err := uc.service.ForgotPassword(resetRequest)
			if err != nil {
				utils.HandleNotFound(ctx, map[string]string{
					"message": "Email tidak ditemukan",
				})
			} else {
				utils.HandleSuccess(ctx, map[string]string{
					"message": "Berhasil reset password, cek email",
				})
			}
		}
	}

}

func (uc *UserController) changePassword(ctx *gin.Context) {
	id := fmt.Sprintf("%v", ctx.MustGet("id"))
	var changePasswordRequest *web.UserChangePasswordRequest

	err := ctx.ShouldBindJSON(&changePasswordRequest)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		uc.service.ChangePassword(id, changePasswordRequest)
		utils.HandleSuccess(ctx, map[string]string{
			"message": "Password berhasil diubah",
		})
	}
}

func (uc *UserController) updateUser(ctx *gin.Context) {
	id := fmt.Sprintf("%v", ctx.MustGet("id"))
	var userUpdateRequest *web.UserUpdateRequest

	err := ctx.ShouldBindJSON(&userUpdateRequest)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := uc.service.UpdateUser(id, userUpdateRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, map[string]string{
				"message": err.Error(),
			})
		} else {
			utils.HandleSuccess(ctx, map[string]string{
				"message": "User berhasil diupdate",
			})
		}
	}

}

func NewUserController(r *gin.Engine, service service.UserService, authMdw middleware.AuthMiddleware) *UserController {
	controller := UserController{
		router:  r,
		service: service,
	}

	r.Use(gin.Recovery())
	userRouteGroup := controller.router.Group("/api/user")
	userRouteGroup.GET("/", authMdw.RequireToken(), controller.getUser)
	userRouteGroup.POST("/register", controller.registerUser)
	userRouteGroup.POST("/login", controller.loginUser)
	userRouteGroup.POST("/forgot-password", authMdw.RequireToken(), controller.forgotPassword)
	userRouteGroup.PUT("/change-password", authMdw.RequireToken(), controller.changePassword)
	userRouteGroup.PUT("/update", authMdw.RequireToken(), controller.updateUser)

	return &controller
}
