package account

import (
	"context"

	"github.com/khdoba2000/banking/configs"
	"github.com/khdoba2000/banking/entities"
	"github.com/khdoba2000/banking/logger"
	"github.com/khdoba2000/banking/storage"
)

// AccountController
type AccountController interface {
	// ListByOwnerID list of accounts of a customer by ID
	GetByOwnerID(ctx context.Context, ownerID string) (*entities.Account, error)
}

type accountController struct {
	log     logger.LoggerI
	storage storage.Storage
	cfg     *configs.Configuration
}

// NewAccountController ...
func NewAccountController(log logger.LoggerI, storage storage.Storage) AccountController {
	return accountController{
		log:     log,
		storage: storage,
		cfg:     configs.Config(),
	}
}

func (c accountController) GetByOwnerID(ctx context.Context, ownerID string) (*entities.Account, error) {

	accounts, err := c.storage.Account().GetByOwnerID(ctx, ownerID)
	if err != nil {
		c.log.Error("calling GetByOwnerID failed", logger.Error(err))
		return nil, err
	}

	return accounts, nil
}
