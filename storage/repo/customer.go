package repo

import (
	"context"

	"github.com/khdoba/banking/entities"
)

// ICustomerStorage customer storage interface
type ICustomerStorage interface {
	Create(ctx context.Context, req entities.Customer) error
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entities.Customer, error)
}
