package repo

import (
	"context"

	"github.com/khdoba/banking/entities"
)

// Defining Base interface for Authentication
type IAuthStorage interface {
	// Login(context.Context, entities.LoginReq) (entities.LoginRes, error)
	Signup(ctx context.Context, req entities.SignupReq) (*entities.SignupRes, error)
}

type ICustomerStorage interface {
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entities.Customer, error)
}
