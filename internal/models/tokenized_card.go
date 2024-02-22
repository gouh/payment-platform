package models

import "github.com/google/uuid"

type TokenizedCard struct {
	Token          string    `db:"token"`
	LastFourDigits string    `db:"last_four_digits"`
	ExpiryMonth    int       `db:"expiry_month"`
	ExpiryYear     int       `db:"expiry_year"`
	CardType       string    `db:"card_type"`
	CustomerID     uuid.UUID `db:"customer_id"`
}
