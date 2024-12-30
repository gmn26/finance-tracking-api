package repository

import (
	"errors"
	"log"

	"github.com/gmn26/finance-tracking-api/model"
	"gorm.io/gorm"
)

type TransactionRepositoryImplType struct {
	Db *gorm.DB
}

func TransactionRepositoryImpl(Db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImplType{Db: Db}
}

func (t TransactionRepositoryImplType) Create(transaction model.Transaction) {
	result := t.Db.Create(&transaction)
	log.Fatal(result.Error)
}

// func (t TransactionRepositoryImplType) Update(transaction mdoel.Transaction) {
// 	var updateTransaction = request.UpdateTransactionRequest{
// 		Id: transaction.id,

// 	}
// }

func (t TransactionRepositoryImplType) Delete(transactionId int) {
	var transaction model.Transaction
	result := t.Db.Where("id = ?", transactionId).Delete(&transaction)
	log.Fatal(result.Error)
}

func (t TransactionRepositoryImplType) FindById(transactionId int) (model.Transaction, error) {
	var transaction model.Transaction
	result := t.Db.Find(&transaction, transactionId)
	if result != nil {
		return transaction, nil
	} else {
		return transaction, errors.New("no transaction")
	}
}

func (t TransactionRepositoryImplType) FindAll() []model.Transaction {
	var transaction []model.Transaction
	results := t.Db.Find(&transaction)
	log.Fatal(results.Error)
	return transaction
}
