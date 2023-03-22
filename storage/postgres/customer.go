package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khdoba2000/banking/constants"
	"github.com/khdoba2000/banking/entities"
	e "github.com/khdoba2000/banking/errors"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type customerRepo struct {
	db *sqlx.DB
}

// NewCustomer postgres implementation of customer storage interface
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

// Create
func (r *customerRepo) Create(ctx context.Context, req entities.Customer) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.ExecContext(ctx, `
	      INSERT INTO customers (id, name, phone_number, password, created_at)
		  VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
	`, req.ID, req.Name, req.PhoneNumber, req.Password)

	if err != nil {
		pgErr, isPGErr := err.(*pq.Error)
		if isPGErr {
			if pgErr.Code == constants.PGUniqueKeyViolationCode {
				return e.ErrCustomerAlreadyExists
			}
		}
		return err
	}

	return nil
}
