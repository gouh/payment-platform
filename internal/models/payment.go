package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID              uuid.UUID `db:"id" json:"id"`
	MerchantID      uuid.UUID `db:"merchant_id" json:"merchantID"`
	Token           string    `db:"token" json:"token"`
	Amount          float64   `db:"amount" json:"amount"`
	Status          string    `db:"status" json:"status"`
	TransactionDate time.Time `db:"transaction_date" json:"transactionDate"`
}
