package main

import "errors"

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

func updateBalance(customer *customer, tx transaction) error {
	if customer.id != tx.customerID {
		return errors.New("this transaction belongs to other customer")
	}

	switch tx.transactionType {
	case transactionDeposit:
		customer.balance += tx.amount
	case transactionWithdrawal:
		if customer.balance < tx.amount {
			return errors.New("insufficient funds")
		}
		customer.balance -= tx.amount
	default:
		return errors.New("unknown transaction type")
	}

	return nil
}
