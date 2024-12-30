package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	Income  TransactionType = "income"
	Outcome TransactionType = "outcome"
)

type Transaction struct {
	Id              uuid.UUID       `gorm:"type:uuid;primaryKey"`
	TransactionName string          `gorm:"type:varchar(255);not null"`
	TransactionType TransactionType `gorm:"type:enum('income','outcome');not null"`
	TransactionTime time.Time       `gorm:"not null"`
}
