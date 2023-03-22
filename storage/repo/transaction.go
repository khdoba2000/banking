package repo

import (
	"context"

	"github.com/khdoba2000/banking/entities"
)

// ITransactionStorage transaction storage interface
type ITransactionStorage interface {
	Create(ctx context.Context, req entities.Transaction) (*entities.TransactionOut, error)
}
