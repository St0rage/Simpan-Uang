package controller

import (
	"fmt"
	"strconv"

	"github.com/St0rage/Simpan-Uang/delivery/middleware"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/service"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/gin-gonic/gin"
)

type WishlistController struct {
	router               *gin.Engine
	wishlistService      service.WishlistService
	wishlistTransService service.WishlistTransactionService
}

func (wc *WishlistController) GetWishlist(ctx *gin.Context) {
	Userid := fmt.Sprintf("%v", ctx.MustGet("id"))

	wishlists := wc.wishlistService.GetWishlist(Userid)

	utils.HandleSuccess(ctx, "Berhasil get wishlist", wishlists)
}

func (wc *WishlistController) GetWishlistById(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlistId")

	wishlistRespone := wc.wishlistService.GetWishlistById(wishlistId)

	utils.HandleSuccess(ctx, "Berhasil get detail wishlist", wishlistRespone)
}

func (wc *WishlistController) CreateNewWishlist(ctx *gin.Context) {
	Userid := fmt.Sprintf("%v", ctx.MustGet("id"))
	var wishlist *web.WishlistCreateUpdateRequest
	err := ctx.ShouldBindJSON(&wishlist)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		getError, err := wc.wishlistService.CreateNewWishlist(Userid, wishlist)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation error", getError)
		} else {
			utils.HandleSuccessCreated(ctx, "Wishlist berhasil dibuat", nil)
		}
	}
}

func (wc *WishlistController) UpdateWishlist(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlistId")
	var wishlistUpdate *web.WishlistCreateUpdateRequest

	err := ctx.ShouldBindJSON(&wishlistUpdate)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		getError, err := wc.wishlistService.UpdateWishlist(wishlistId, wishlistUpdate)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation error", getError)
		} else {
			utils.HandleSuccess(ctx, "Wishlist berhasil diupdate", nil)
		}
	}
}

func (wc *WishlistController) DeleteWishlist(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	wishlistId := ctx.Param("wishlistId")

	wc.wishlistService.DeleteWishlist(userId, wishlistId)
	utils.HandleSuccess(ctx, "Wishlist Berhasil dihapus", nil)

}

func (wc *WishlistController) DepositWishlist(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlistId")
	wishlistTarget := wc.wishlistService.GetWishlistTarget(wishlistId)
	var depositTransaction *web.DepositTransactionRequest

	err := ctx.ShouldBindJSON(&depositTransaction)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation error", getError)
	} else {
		getError, err := wc.wishlistTransService.DepositWishlist(wishlistId, wishlistTarget, depositTransaction)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation Error", getError)
		} else {
			utils.HandleSuccessCreated(ctx, "Transaksi Sebesar Rp "+strconv.Itoa(int(depositTransaction.Amount.(float64)))+" Berhasil Masuk", nil)
		}
	}
}

func (wc *WishlistController) WithdrawWishlist(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlistId")
	var withdrawTransaction *web.WithdrawTransactionRequest

	err := ctx.ShouldBindJSON(&withdrawTransaction)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation Error", getError)
	} else {
		getError, err := wc.wishlistTransService.WithdrawWishlist(wishlistId, withdrawTransaction)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation Error", getError)
		} else {
			utils.HandleSuccessCreated(ctx, "Transaksi Sebesar Rp "+strconv.Itoa(int(withdrawTransaction.Amount.(float64)))+" Berhasil ditarik", nil)
		}
	}
}

func (wc *WishlistController) GetWishlistTransactions(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlistId")
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		page = 1
	}

	transactions := wc.wishlistTransService.GetWishlistTransaction(wishlistId, page)

	utils.HandleSuccess(ctx, "Berhasil get transaksi wishlist", transactions)
}

func (wc *WishlistController) DeleteWishlistTransactions(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlistId")
	wishlistTransId := ctx.Param("wishlistTransId")

	err := wc.wishlistTransService.DeleteTransaction(wishlistTransId, wishlistId)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error(), nil)
	} else {
		utils.HandleSuccess(ctx, "Transaksi berhasil dihapus", nil)
	}
}

func NewWishlistController(r *gin.Engine, wishlistService service.WishlistService, wishlistTransService service.WishlistTransactionService, authMdw middleware.AuthMiddleware) *WishlistController {
	controller := WishlistController{
		router:               r,
		wishlistService:      wishlistService,
		wishlistTransService: wishlistTransService,
	}

	wishlistRouteGroup := controller.router.Group("/api/wishlist", authMdw.RequireToken())
	wishlistRouteGroup.GET("/", controller.GetWishlist)
	wishlistRouteGroup.GET("/:wishlistId", authMdw.WishlistAuthorization(), controller.GetWishlistById)
	wishlistRouteGroup.POST("/create", controller.CreateNewWishlist)
	wishlistRouteGroup.PUT("/:wishlistId/update", authMdw.WishlistAuthorization(), controller.UpdateWishlist)
	wishlistRouteGroup.DELETE("/:wishlistId/delete", authMdw.WishlistAuthorization(), controller.DeleteWishlist)
	wishlistRouteGroup.GET("/:wishlistId/transactions", authMdw.WishlistAuthorization(), controller.GetWishlistTransactions)
	wishlistRouteGroup.POST("/:wishlistId/transactions/deposit", authMdw.WishlistAuthorization(), controller.DepositWishlist)
	wishlistRouteGroup.POST("/:wishlistId/transactions/withdraw", authMdw.WishlistAuthorization(), controller.WithdrawWishlist)
	wishlistRouteGroup.DELETE("/:wishlistId/transactions/:wishlistTransId/delete", authMdw.WishlistAuthorization(), controller.DeleteWishlistTransactions)

	return &controller
}
