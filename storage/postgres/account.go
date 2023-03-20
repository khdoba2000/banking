package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khdoba/banking/constants"
	"github.com/khdoba/banking/entities"
	e "github.com/khdoba/banking/errors"
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
			err = tx.Rollback()
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

// ListByOwnerID
func (r *accountRepo) ListByOwnerID(ctx context.Context, ownerID string) ([]entities.Account, error) {
	rows, err := r.db.Query(`
		SELECT id, currency_code, balance
		FROM accounts
		WHERE 
			owner_id = $1
	`, ownerID)

	if err != nil {
		fmt.Println("QueryRow  error:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, e.ErrAccountNotExists
		}
		return nil, err
	}

	var accounts []entities.Account
	for rows.Next() {
		account := entities.Account{}
		err := rows.Scan(
			&account.ID,
			&account.CurrencyCode,
			&account.Balance,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
