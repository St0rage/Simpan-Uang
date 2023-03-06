package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWrapper struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func HandleSuccess(ctx *gin.Context, message string, data interface{}) {
	response := ResponseWrapper{
		Code:    200,
		Status:  "OK",
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func HandleSuccessCreated(ctx *gin.Context, message string, data interface{}) {
	response := ResponseWrapper{
		Code:    201,
		Status:  "CREATED",
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusCreated, response)
}

func HandleNotFound(ctx *gin.Context, data interface{}) {
	response := ResponseWrapper{
		Code:   404,
		Status: "NOT FOUND",
	}
	ctx.JSON(http.StatusNotFound, response)
}

func HandleInternalServerError(ctx *gin.Context) {
	response := ResponseWrapper{
		Code:   500,
		Status: "INTERNAL SERVER ERROR",
	}
	ctx.JSON(http.StatusInternalServerError, response)
}

func HandleBadRequest(ctx *gin.Context, message string, errors interface{}) {
	response := ResponseWrapper{
		Code:    400,
		Status:  "BAD REQUEST",
		Message: message,
		Errors:  errors,
	}
	ctx.JSON(http.StatusBadRequest, response)
}

func HandleUnauthorized(ctx *gin.Context) {
	response := ResponseWrapper{
		Code:   401,
		Status: "UNAUTHORIZED",
	}
	ctx.JSON(http.StatusUnauthorized, response)
}
