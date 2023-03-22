package transaction

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/khdoba2000/banking/configs"
	"github.com/khdoba2000/banking/entities"
	"github.com/khdoba2000/banking/logger"
	e "github.com/khdoba2000/banking/pkg/errors"
	"github.com/khdoba2000/banking/storage"
)

// TransactionController
type TransactionController interface {
	Create(ctx context.Context, req entities.CreateTransactionReq) (*entities.TransactionOut, error)
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
func (tc transactionController) Create(ctx context.Context, req entities.CreateTransactionReq) (*entities.TransactionOut, error) {

	if req.Transaction.IsOut() {
		// check if the balance is sufficient to do the transaction
		account, err := tc.storage.Account().GetByOwnerID(ctx, req.CustomerID)
		if err != nil {
			tc.log.Error("calling Account.GetByOwnerID failed", logger.Error(err))
			return nil, err
		}
		if account.Balance < req.Transaction.GetAmount() {
			tc.log.Error("unsufficient balance")
			return nil, e.NewError(http.StatusBadRequest, "unsufficient balance")
		}
	}

	req.Transaction.SetID(uuid.NewString())

	res, err := tc.storage.Transaction().Create(ctx, req.Transaction)
	if err != nil {
		tc.log.Error("calling Transaction.Create failed", logger.Error(err))
		return nil, err
	}

	return res, nil
}
