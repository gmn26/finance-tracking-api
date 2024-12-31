package service

import (
	"log"

	"github.com/gmn26/finance-tracking-api/data/request"
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

func (t TransactionServiceImplType) Create(transaction request.CreateTransactionRequest) {
	err := t.Validate.Struct(transaction)
	if err != nil {
		log.Fatal(err)
	}
	transactionModel := model.Transaction{
		TransactionName: transaction.TransactionName,
	}
	t.TransactionRepository.Create(transactionModel)
}
