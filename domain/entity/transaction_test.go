package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionIsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 999
	assert.Nil(t, transaction.isValid())
}

func TestTransactionIsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001

	err := transaction.isValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransactionIsNotValidWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	err := transaction.isValid()
	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
