package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID              uuid.UUID `db:"id"`
	MerchantID      uuid.UUID `db:"merchant_id"`
	Token           string    `db:"token"`
	Amount          float64   `db:"amount"`
	Status          string    `db:"status"`
	TransactionDate time.Time `db:"transaction_date"`
}
