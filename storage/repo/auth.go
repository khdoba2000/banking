package repo

import (
	"context"

	"github.com/khdoba/banking/entities"
)

// Defining Base interface for Authentication
type IAuthStorage interface {
	Login(context.Context, entities.LoginReq) (entities.LoginRes, error)
	Signup(ctx context.Context, data entities.SignupReq) (entities.SignupRes, error)
}
