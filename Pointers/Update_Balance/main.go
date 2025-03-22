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

// ?

func updateBalance(c *customer, t transaction) error {
	if t.transactionType != transactionDeposit && t.transactionType != transactionWithdrawal {
		return errors.New("unknown transaction type")
	}

	if t.transactionType == transactionDeposit {
		c.balance += t.amount
		return nil
	} else if c.balance >= t.amount{
		c.balance -= t.amount
		return nil
	}else {
		return errors.New("insufficient funds")
	}

}