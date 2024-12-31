package request

import (
	"time"
	"github.com/gmn26/finance-tracking-api/model"
)

type CreateTransactionRequest struct {
	TransactionName string `validate:"required,min=1,max=255"`
	TransactionType model.TransactionType `validate:"required,oneof=income outcome"`
	TransactionTime time.Time `validate:"required"`
}