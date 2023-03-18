package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khdoba/banking/entities"
	e "github.com/khdoba/banking/errors"
	_ "github.com/lib/pq"
)

type customerRepo struct {
	db *sqlx.DB
}

// NewCustomer ...
func NewCustomer(db *sqlx.DB) *customerRepo {
	return &customerRepo{db: db}
}

// GetByPhoneNumber
func (r *customerRepo) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entities.Customer, error) {
	row := r.db.QueryRow(`
		SELECT id, name, phone_number, password
		FROM customers
		WHERE 
			phone_number = $1
	`, phoneNumber)

	if row.Err() != nil {
		fmt.Println("QueryRow  error:", row.Err())
		return nil, row.Err()
	}

	customer := entities.Customer{}
	err := row.Scan(&customer.ID,
		&customer.Name,
		&customer.PhoneNumber,
		&customer.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, e.ErrCustomerNotExists
		}
		return nil, err
	}

	return &customer, nil
}
