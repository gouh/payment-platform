package models

import "github.com/google/uuid"

type TokenizedCard struct {
	Token          string    `db:"token" json:"token"`
	LastFourDigits string    `db:"last_four_digits" json:"lastFourDigits"`
	ExpiryMonth    int       `db:"expiry_month" json:"expiryMonth"`
	ExpiryYear     int       `db:"expiry_year" json:"expiryYear"`
	CardType       string    `db:"card_type" json:"cardType"`
	CustomerID     uuid.UUID `db:"customer_id" json:"customerID"`
}
