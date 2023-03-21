package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/entities"
	"github.com/khdoba/banking/logger"
	"github.com/khdoba/banking/storage"
)

// TransactionController
type TransactionController interface {
	Create(ctx context.Context, req entities.Transaction) (*entities.TransactionOut, error)
}

type transactionController struct {
	log     logger.LoggerI
	storage storage.Storage
	cfg     *configs.Configuration
}

// NewTransactionController ...
func NewTransactionController(log logger.LoggerI, storage storage.Storage) TransactionController {
	return transactionController{
		log:     log,
		storage: storage,
		cfg:     configs.Config(),
	}
}

//

// Create ...
func (tc transactionController) Create(ctx context.Context, req entities.Transaction) (*entities.TransactionOut, error) {

	req.SetID(uuid.NewString())

	res, err := tc.storage.Transaction().Create(ctx, req)
	if err != nil {
		tc.log.Error("calling Transaction.Create failed", logger.Error(err))
		return nil, err
	}

	return res, nil
}
