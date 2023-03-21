package repo

import (
	"context"

	"github.com/khdoba/banking/entities"
)

// IAccountStorage account storage interface
type IAccountStorage interface {
	Create(ctx context.Context, req entities.CreateAccountReq) error
	GetByOwnerID(ctx context.Context, ownerID string) (*entities.Account, error)
}
