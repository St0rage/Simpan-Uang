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

type PiggyBankController struct {
	router                *gin.Engine
	piggyBankService      service.PiggyBankService
	piggyBankTransService service.PiggyBankTransactionService
}

// PiggyBank
func (pc *PiggyBankController) CreatePiggyBank(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	var newPiggyBank *web.PiggyBankCreateUpdateRequest
	err := ctx.ShouldBindJSON(&newPiggyBank)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation Error", getError)
	} else {
		getError, err := pc.piggyBankService.CreatePiggyBank(userId, newPiggyBank)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation Error", getError)
		} else {
			utils.HandleSuccessCreated(ctx, "Tabungan berhasil dibuat", nil)
		}
	}
}

func (pc *PiggyBankController) GetPiggyBanks(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	piggyBankReponses := pc.piggyBankService.GetAllPiggyBank(userId)

	utils.HandleSuccess(ctx, "Berhasil get tabungan", piggyBankReponses)
}

func (pc *PiggyBankController) GetPiggyBankById(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")

	piggyBankRespone := pc.piggyBankService.GetPiggyBankById(piggyBankId)

	utils.HandleSuccess(ctx, "Berhasil get detail tabungan", piggyBankRespone)
}

func (pc *PiggyBankController) UpdatePiggyBank(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")
	var piggyBankUpdate *web.PiggyBankCreateUpdateRequest

	err := ctx.ShouldBindJSON(&piggyBankUpdate)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation Error", getError)
	} else {
		getError, err := pc.piggyBankService.UpdatePiggyBank(piggyBankId, piggyBankUpdate)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation Error", getError)
		} else {
			utils.HandleSuccess(ctx, "Tabungan berhasil diupdate", nil)
		}
	}
}

func (pc *PiggyBankController) DeletePiggyBank(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	piggyBankId := ctx.Param("piggyBankId")

	err := pc.piggyBankService.DeletePiggyBank(userId, piggyBankId)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error(), nil)
	} else {
		utils.HandleSuccess(ctx, "Tabungan Berhasil dihapus", nil)
	}

}

// PiggyBankTransaction
func (pc *PiggyBankController) DepositPiggyBank(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")
	var depositTransaction *web.DepositTransactionRequest

	err := ctx.ShouldBindJSON(&depositTransaction)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation Error", getError)
	} else {
		pc.piggyBankTransService.DepositTransaction(piggyBankId, depositTransaction)
		utils.HandleSuccessCreated(ctx, "Transaksi Sebesar Rp "+strconv.Itoa(int(depositTransaction.Amount.(float64)))+" Berhasil Masuk", nil)
	}
}

func (pc *PiggyBankController) WithdrawPiggyBank(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")
	var withdrawTransaction *web.WithdrawTransactionRequest

	err := ctx.ShouldBindJSON(&withdrawTransaction)
	if err != nil {
		getError := utils.CustomValidationErr(err)
		utils.HandleBadRequest(ctx, "Validation Error", getError)
	} else {
		getError, err := pc.piggyBankTransService.WithdrawTransaction(piggyBankId, withdrawTransaction)
		if err != nil {
			utils.HandleBadRequest(ctx, "Validation Error", getError)
		} else {
			utils.HandleSuccessCreated(ctx, "Transaksi Sebesar Rp "+strconv.Itoa(int(withdrawTransaction.Amount.(float64)))+" Berhasil ditarik", nil)
		}
	}
}

func (pc *PiggyBankController) GetPiggyBankTransactions(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		page = 1
	}

	transactions := pc.piggyBankTransService.GetAllTransactions(piggyBankId, page)

	utils.HandleSuccess(ctx, "Berhasil get transaksi tabungan", transactions)
}

func (pc *PiggyBankController) DeletePiggyBankTransactions(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")
	piggyBankTransId := ctx.Param("piggyBankTransId")

	err := pc.piggyBankTransService.DeleteTransaction(piggyBankTransId, piggyBankId)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error(), nil)
	} else {
		utils.HandleSuccess(ctx, "Transaksi berhasil dihapus", nil)
	}
}

func NewPiggyBankController(r *gin.Engine, piggyBankService service.PiggyBankService, piggyBankTransService service.PiggyBankTransactionService, authMdw middleware.AuthMiddleware) *PiggyBankController {
	controller := PiggyBankController{
		router:                r,
		piggyBankService:      piggyBankService,
		piggyBankTransService: piggyBankTransService,
	}

	controller.router.Use(gin.Recovery())
	piggyBankRouteGroup := controller.router.Group("/api/piggy-bank", authMdw.RequireToken())
	// piggy-bank
	piggyBankRouteGroup.GET("/", controller.GetPiggyBanks)
	piggyBankRouteGroup.GET("/:piggyBankId", authMdw.PiggyBankAuthorization(), controller.GetPiggyBankById)
	piggyBankRouteGroup.POST("/create", controller.CreatePiggyBank)
	piggyBankRouteGroup.PUT("/:piggyBankId/update", authMdw.PiggyBankAuthorization(), controller.UpdatePiggyBank)
	piggyBankRouteGroup.DELETE("/:piggyBankId/delete", authMdw.PiggyBankAuthorization(), controller.DeletePiggyBank)
	// piggy-bank-transaction
	piggyBankRouteGroup.GET("/:piggyBankId/transactions", authMdw.PiggyBankAuthorization(), controller.GetPiggyBankTransactions)
	piggyBankRouteGroup.POST("/:piggyBankId/transactions/deposit", authMdw.PiggyBankAuthorization(), controller.DepositPiggyBank)
	piggyBankRouteGroup.POST("/:piggyBankId/transactions/withdraw", authMdw.PiggyBankAuthorization(), controller.WithdrawPiggyBank)
	piggyBankRouteGroup.DELETE("/:piggyBankId/transactions/:piggyBankTransId/delete", authMdw.PiggyBankAuthorization(), controller.DeletePiggyBankTransactions)

	return &controller
}
