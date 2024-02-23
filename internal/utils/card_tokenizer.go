package utils

import (
	"braces.dev/errtrace"
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
)

func TokenizeCard(card requests.TokenizedCardRequest, customerId string) (*models.TokenizedCard, error) {
	key := fmt.Sprintf("%s-%d-%d-%s-%s", card.CardNumber, card.ExpiryMonth, card.ExpiryYear, card.CardType, customerId)

	hash := sha256.Sum256([]byte(key))
	token := fmt.Sprintf("%x", hash[:])

	uuidCustomer, errUuid := uuid.Parse(customerId)
	if errUuid != nil {
		return nil, errtrace.Wrap(errUuid)
	}

	tokenizedCard := models.TokenizedCard{
		Token:          token,
		LastFourDigits: card.CardNumber[len(card.CardNumber)-4:],
		ExpiryMonth:    card.ExpiryMonth,
		ExpiryYear:     card.ExpiryYear,
		CardType:       card.CardType,
		CustomerID:     uuidCustomer,
	}

	return &tokenizedCard, nil
}
