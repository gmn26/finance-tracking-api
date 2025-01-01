package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gmn26/finance-tracking-api/controller"
)

func TransactionRouter(transactionCOntroller *controller.TransactionController) *gin.Engine {
	service := gin.Default()

	transactionRouter := service.Group("/api/transaction/")

	// transactionRouter.GET("/", transactionCOntroller)
	transactionRouter.POST("/", transactionCOntroller.Create)

	return service
}
