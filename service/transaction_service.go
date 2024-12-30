package service

type TransactionService interface {
	// Create param type request.CreateTransactionRequest
	Create(transaction any)
	//Delete(transactionId int)
	//FindById(transactionId int) response.TransactionResponse
	//FindAll() []response.TransactionResponse
}
