package controller

import (
	"fmt"

	"github.com/St0rage/Simpan-Uang/delivery/middleware"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/service"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/gin-gonic/gin"
)

type PiggyBankController struct {
	router  *gin.Engine
	service service.PiggyBankService
}

func (pc *PiggyBankController) CreatePiggyBank(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	var newPiggyBank *web.PiggyBankCreateUpdateRequest
	err := ctx.ShouldBindJSON(&newPiggyBank)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := pc.service.CreatePiggyBank(userId, newPiggyBank)
		if err != nil {
			utils.HandleBadRequest(ctx, map[string]string{
				"message": err.Error(),
			})
		} else {
			utils.HandleSuccessCreated(ctx, map[string]string{
				"message": "Tabungan berhasil dibuat",
			})
		}
	}
}

func (pc *PiggyBankController) GetPiggyBanks(ctx *gin.Context) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	piggyBankReponses := pc.service.GetAllPiggyBank(userId)

	utils.HandleSuccess(ctx, piggyBankReponses)
}

func (pc *PiggyBankController) GetPiggyBankById(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")

	piggyBankRespone := pc.service.GetPiggyBankById(piggyBankId)

	utils.HandleSuccess(ctx, piggyBankRespone)
}

func (pc *PiggyBankController) UpdatePiggyBank(ctx *gin.Context) {
	piggyBankId := ctx.Param("piggyBankId")
	var piggyBankUpdate *web.PiggyBankCreateUpdateRequest

	err := ctx.ShouldBindJSON(&piggyBankUpdate)
	if err != nil {
		utils.HandleBadRequest(ctx, err.Error())
	} else {
		err := pc.service.UpdatePiggyBank(piggyBankId, piggyBankUpdate)
		if err != nil {
			utils.HandleBadRequest(ctx, map[string]string{
				"message": err.Error(),
			})
		} else {
			utils.HandleSuccess(ctx, map[string]string{
				"message": "Tabungan berhasil diupdate",
			})
		}
	}
}

func NewPiggyBankController(r *gin.Engine, service service.PiggyBankService, authMdw middleware.AuthMiddleware) *PiggyBankController {
	controller := PiggyBankController{
		router:  r,
		service: service,
	}

	controller.router.Use(gin.Recovery())
	piggyBankRouteGroup := controller.router.Group("/api/piggy-bank", authMdw.RequireToken())
	piggyBankRouteGroup.GET("/", controller.GetPiggyBanks)
	piggyBankRouteGroup.GET("/:piggyBankId", authMdw.PiggyBankAuthorization(), controller.GetPiggyBankById)
	piggyBankRouteGroup.POST("/create", controller.CreatePiggyBank)
	piggyBankRouteGroup.PUT("/:piggyBankId/update", authMdw.PiggyBankAuthorization(), controller.UpdatePiggyBank)

	return &controller
}
