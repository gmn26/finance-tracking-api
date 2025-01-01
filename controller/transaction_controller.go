package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmn26/finance-tracking-api/data/request"
	"github.com/gmn26/finance-tracking-api/data/response"
	"github.com/gmn26/finance-tracking-api/service"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(service service.TransactionService) *TransactionController {
	return &TransactionController{transactionService: service}
}

func (controller *TransactionController) Create(ctx *gin.Context) {
	createTransactionRequest := request.CreateTransactionRequest{}
	err := ctx.ShouldBindJSON(&createTransactionRequest)
	if err != nil {
		log.Fatal(err)
	}

	controller.transactionService.Create(createTransactionRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
