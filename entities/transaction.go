package entities

import (
	"time"
)

type Transaction interface {
	Entity
	Validator
	GetAmount() uint64
	GetAccountID() string
	GetTypeID() int
	IsOut() bool
}

type TransactionImp struct {
	ID        string
	Amount    uint64
	AccountID string
}

// GetAmount ...
func (t *TransactionImp) GetAmount() uint64 {
	return t.Amount
}

// GetID ...
func (t *TransactionImp) GetID() string {
	return t.ID
}

// SetID ...
func (t *TransactionImp) SetID(id string) {
	t.ID = id
}

// GetAccountID ...
func (t *TransactionImp) GetAccountID() string {
	return t.AccountID
}

type TransactionOut struct {
	ID            string
	TypeName      string
	AccountFromID string
	AccountToID   string
	Amount        uint64
	CreatedAt     time.Time
}

type CreateTransactionReq struct {
	Transaction Transaction
	CustomerID  string
}
