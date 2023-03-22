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
)

type accountRepo struct {
	db *sqlx.DB
}

// NewAccount postgres implementation of account storage interface
func NewAccount(db *sqlx.DB) *accountRepo {
	return &accountRepo{db: db}
}

// Create
func (r *accountRepo) Create(ctx context.Context, req entities.CreateAccountReq) error {

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
	      INSERT INTO accounts (id, owner_id, created_at)
		  VALUES ($1, $2, CURRENT_TIMESTAMP)
	`, req.ID, req.OwnerID)

	if err != nil {
		pgErr, isPGErr := err.(*pq.Error)
		if isPGErr {
			if pgErr.Code == constants.PGUniqueKeyViolationCode {
				return e.ErrAccountAlreadyExists
			}
		}
		return err
	}

	return nil
}

// GetByOwnerID
func (r *accountRepo) GetByOwnerID(ctx context.Context, ownerID string) (*entities.Account, error) {
	row := r.db.QueryRow(`
		SELECT id, balance
		FROM accounts
		WHERE 
			owner_id = $1
			AND currency_code = 'UZS'
		LIMIT 1
	`, ownerID)

	if row.Err() != nil {
		fmt.Println("QueryRow  error:", row.Err())
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, e.ErrAccountNotExists
		}
		return nil, row.Err()
	}

	account := entities.Account{}
	err := row.Scan(
		&account.ID,
		&account.Balance,
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
