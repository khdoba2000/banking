package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/khdoba/banking/entities"
	_ "github.com/lib/pq"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *authRepo {
	return &authRepo{db: db}
}

// Login
func (r *authRepo) Login(ctx context.Context, req entities.LoginReq) (entities.LoginRes, error) {
	return entities.LoginRes{}, nil
}

// Signup
func (r *authRepo) Signup(ctx context.Context, req entities.SignupReq) (entities.SignupRes, error) {
	return entities.SignupRes{}, nil
}
