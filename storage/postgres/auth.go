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

// // Login
// func (r *authRepo) Login(ctx context.Context, req entities.LoginReq) (entities.LoginRes, error) {
// 	res:=r.db.QueryRow(`SELECT id, name
// 				 FROM customers
// 				WHERE
// 				 phone_number = $1
// 				 	AND
// 				 password = $2
// 	`, req.PhoneNumber, req.Password)

// 	if res.Err()
// 	res.Scan()

// 	return entities.LoginRes{}, nil
// }

// Signup
func (r *authRepo) Signup(ctx context.Context, req entities.SignupReq) (*entities.SignupRes, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.ExecContext(ctx, `
	      INSERT INTO customers (id, name, phone_number, password, created_at)
		  VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
	`, req.ID, req.Name, req.PhoneNumber, req.Password)

	if err != nil {
		return nil, err
	}

	return &entities.SignupRes{ID: req.ID}, nil
}
