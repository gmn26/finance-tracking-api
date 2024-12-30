package service

import (
	"log"

	"github.com/gmn26/finance-tracking-api/model"
	"github.com/gmn26/finance-tracking-api/repository"
	"github.com/go-playground/validator/v10"
)

type TransactionServiceImplType struct {
	TransactionRepository repository.TransactionRepository
	Validate              *validator.Validate
}

func TransactionRepositoryImpl(transactionRepository repository.TransactionRepository, validate *validator.Validate) TransactionService {
	return &TransactionServiceImplType{
		TransactionRepository: transactionRepository,
		Validate:              validate,
	}
}

// Create param type request.CreateTransactionRequest
func (t TransactionServiceImplType) Create(transaction any) {
	err := t.Validate.Struct(transaction)
	log.Fatal(err)
	transactionModel := model.Transaction{
		TransactionName: "transaction.name",
	}
	t.TransactionRepository.Create(transactionModel)
}
