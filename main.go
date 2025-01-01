package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmn26/finance-tracking-api/config"
	"github.com/gmn26/finance-tracking-api/controller"
	"github.com/gmn26/finance-tracking-api/model"
	"github.com/gmn26/finance-tracking-api/repository"
	"github.com/gmn26/finance-tracking-api/router"
	"github.com/gmn26/finance-tracking-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("transactions").AutoMigrate(&model.Transaction{})

	transactionRepository := repository.TransactionRepositoryImpl(db)

	transactionService := service.TransactionRepositoryImpl(transactionRepository, validate)

	transactionController := controller.NewTransactionController(transactionService)

	routes := router.TransactionRouter(transactionController)

	server := &http.Server{
		Addr:           ":6969",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
