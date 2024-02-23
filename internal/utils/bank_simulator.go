package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type (
	BankSimulatorInterface interface {
		AuthorizePayment(string) error
		ValidateCard(card string) error
	}
	BankSimulator struct{}
)

func (bs *BankSimulator) AuthorizePayment(paymentId string) error {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 80% success
	if newRand.Intn(10) < 2 {
		return errors.New(fmt.Sprintf("payment %s authorization failed", paymentId))
	}
	return nil
}

func (bs *BankSimulator) ValidateCard(card string) error {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 90% success
	if newRand.Intn(10) > 1 {
		return errors.New(fmt.Sprintf("card %s is invalid", card))
	}
	return nil
}

func NewBankSimulator() *BankSimulator {
	return &BankSimulator{}
}
