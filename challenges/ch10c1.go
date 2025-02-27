package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(customer *customer, transaction transaction) error {
	if transaction.transactionType != transactionDeposit && transaction.transactionType != transactionWithdrawal {
		return errors.New("unknown transaction type")
	}

	if transaction.transactionType == transactionDeposit {
		customer.balance += transaction.amount
	}

	if transaction.transactionType == transactionWithdrawal {
		if customer.balance < transaction.amount {
			return errors.New("insufficient funds")
		}
		customer.balance -= transaction.amount
	}

	return nil
}
