package utils

import (
	"braces.dev/errtrace"
	"fmt"
	"math/rand"
	"time"
)

type (
	BankSimulatorInterface interface {
		AuthorizePayment(string) error
	}
	BankSimulator struct{}
)

func (bs *BankSimulator) AuthorizePayment(paymentId string) error {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 40% success
	if newRand.Intn(10) < 6 {
		return errtrace.New(fmt.Sprintf("payment %s authorization failed", paymentId))
	}
	return nil
}

func NewBankSimulator() *BankSimulator {
	return &BankSimulator{}
}
