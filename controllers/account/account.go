package account

import (
	"context"
	"errors"
	"net/http"

	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/entities"
	e "github.com/khdoba/banking/errors"
	"github.com/khdoba/banking/logger"
	pkgerrors "github.com/khdoba/banking/pkg/errors"
	"github.com/khdoba/banking/storage"
)

// AccountController
type AccountController interface {
	// ListByOwnerID list of accounts of a customer by ID
	ListByOwnerID(ctx context.Context, ownerID string) ([]entities.Account, error)
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

func (c accountController) ListByOwnerID(ctx context.Context, ownerID string) ([]entities.Account, error) {

	accounts, err := c.storage.Account().ListByOwnerID(ctx, ownerID)
	if err != nil {
		c.log.Error("calling GetByPhoneNumber failed", logger.Error(err))
		if errors.Is(err, e.ErrAccountNotExists) {
			return nil, pkgerrors.NewError(http.StatusForbidden, "phoneNumber or password is wrong")
		}
		return nil, err
	}

	return accounts, nil
}
