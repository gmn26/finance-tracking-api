package controller

import (
	"github.com/gmn26/finance-tracking-api/service"
)

type TransactionController struct {
	transactionService service.TransactionService
}
