package repository

import "github.com/gmn26/finance-tracking-api/model"

type TransactionRepository interface {
	Create(transaction model.Transaction)
	// Update(transaction model.Transaction)
	Delete(transactionId int)
	FindById(transactionId int) (transaction model.Transaction, err error)
	FindAll() []model.Transaction
}
