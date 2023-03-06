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
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	user := uc.service.GetUser(userId)

	utils.HandleSuccess(ctx, "Berhasil get user", user)
}

func (uc *UserController) registerUser(ctx *gin.Context) {
	var newUserRequest *web.UserRegisterRequest
	err := ctx.ShouldBindJSON(&newUserRequest)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		getError, err := uc.service.Register(newUserRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation Error", getError)
		} else {
			utils.HandleSuccessCreated(ctx, "User berhasil dibuat", nil)
		}
	}
}

func (uc *UserController) loginUser(ctx *gin.Context) {
	var loginRequest *web.UserLoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		token, err := uc.service.Login(loginRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, "Email atau Password salah", nil)
		} else {
			utils.HandleSuccess(ctx, "Login Berhasil", gin.H{
				"token": token,
			})
		}
	}
}

func (uc *UserController) forgotPassword(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	isAdmin := uc.service.CheckAdmin(userId)

	if !isAdmin {
		utils.HandleUnauthorized(ctx)
	} else {
		var resetRequest *web.UserResetRequest
		err := ctx.ShouldBindJSON(&resetRequest)
		if err != nil {
			getError := utils.CustomValidationErr(err)
			utils.HandleBadRequest(ctx, "Validation error", getError)
		} else {
			getError, err := uc.service.ForgotPassword(resetRequest)
			if err != nil {
				utils.HandleBadRequest(ctx, "Validation error", getError)
			} else {
				utils.HandleSuccess(ctx, "Berhasil reset password, cek email", nil)
			}
		}
	}

}

func (uc *UserController) changePassword(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var changePasswordRequest *web.UserChangePasswordRequest

	err := ctx.ShouldBindJSON(&changePasswordRequest)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		uc.service.ChangePassword(userId, changePasswordRequest)
		utils.HandleSuccess(ctx, "Password berhasil diubah", nil)
	}
}

func (uc *UserController) updateUser(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var userUpdateRequest *web.UserUpdateRequest

	err := ctx.ShouldBindJSON(&userUpdateRequest)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		getError, err := uc.service.UpdateUser(userId, userUpdateRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation error", getError)
		} else {
			utils.HandleSuccess(ctx, "User berhasil diupdate", nil)
		}
	}
}

func (uc *UserController) updateAvatar(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var avatarUpdateRequest *web.UserAvatarRequest

	err := ctx.ShouldBindJSON(&avatarUpdateRequest)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		getError, err := uc.service.UpdateAvatar(userId, avatarUpdateRequest)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation Error", getError)
		} else {
			utils.HandleSuccess(ctx, "avatar berhasil diupload", nil)
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
	userRouteGroup.PUT("/update-avatar", authMdw.RequireToken(), controller.updateAvatar)

	return &controller
}
