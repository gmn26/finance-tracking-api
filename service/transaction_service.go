package service

import "github.com/gmn26/finance-tracking-api/data/request"

type TransactionService interface {
	// Create param type request.CreateTransactionRequest
	Create(transaction request.CreateTransactionRequest)
	//Delete(transactionId int)
	//FindById(transactionId int) response.TransactionResponse
	//FindAll() []response.TransactionResponse
}
