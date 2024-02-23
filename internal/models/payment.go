package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID              uuid.UUID `db:"id" json:"id"`
	MerchantID      uuid.UUID `db:"merchant_id" json:"merchantId" binding:"required"`
	Token           string    `db:"token" json:"token" binding:"required"`
	Amount          float64   `db:"amount" json:"amount" binding:"required"`
	Status          string    `db:"status" json:"status"`
	TransactionDate time.Time `db:"transaction_date" json:"transactionDate"`
}
