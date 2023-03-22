package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khdoba2000/banking/constants"
	"github.com/khdoba2000/banking/entities"
	e "github.com/khdoba2000/banking/errors"
	"github.com/lib/pq"
)

type transactionRepo struct {
	db *sqlx.DB
}

// NewTransaction postgres implementation of transaction storage interface
func NewTransaction(db *sqlx.DB) *transactionRepo {
	return &transactionRepo{db: db}
}

// Create
func (r *transactionRepo) Create(ctx context.Context, req entities.Transaction) (*entities.TransactionOut, error) {
	var (
		accountFromID sql.NullString
		accountToID   sql.NullString
	)
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			fmt.Println("error:", err)
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	switch transaction := req.(type) {
	case *entities.Income:
		accountToID.Scan(req.GetAccountID())
		// transactionTypeID = constants.IncomeTransactionID
	case *entities.Expense:
		accountFromID.Scan(req.GetAccountID())
		// transactionTypeID = constants.ExpenseTransactionID
	case *entities.Transfer:
		accountFromID.Scan(req.GetAccountID())
		accountToID.Scan(transaction.AccountToID)
		// transactionTypeID = constants.TransferTransactionID
	}

	row := tx.QueryRowContext(ctx, `
	    INSERT INTO transactions (
			id, 
			type_id, 
			account_from_id,
			account_to_id,
			amount,
			created_at
		 )
		 VALUES ($1,
			 $2, 
			 $3,
			 $4,
			 $5,
			 CURRENT_TIMESTAMP)
		 RETURNING 
		 	id, 
			(SELECT tt.name 
				FROM transaction_types tt 
				WHERE tt.id=type_id) AS typeName, 
			account_from_id, 
			account_to_id, 
			amount, 
			created_at
	`, req.GetID(),
		req.GetTypeID(),
		accountFromID,
		accountToID,
		req.GetAmount())

	tr := entities.TransactionOut{}
	err = row.Scan(
		&tr.ID,
		&tr.TypeName,
		&accountFromID,
		&accountToID,
		&tr.Amount,
		&tr.CreatedAt,
	)
	if err != nil {
		pgErr, isPGErr := err.(*pq.Error)
		if isPGErr {
			if pgErr.Code == constants.PGForeignKeyViolationCode {
				return nil, e.ErrAccountNotExists
			}
		}
		return nil, err
	}

	tr.AccountFromID = accountFromID.String
	tr.AccountToID = accountToID.String

	return &tr, nil
}

// // ListByOwnerID
// func (r *transactionRepo) ListByOwnerID(ctx context.Context, ownerID string) ([]entities.Account, error) {
// 	rows, err := r.db.Query(`
// 		SELECT id, currency_code, balance
// 		FROM transactions
// 		WHERE
// 			owner_id = $1
// 	`, ownerID)

// 	if err != nil {
// 		fmt.Println("QueryRow  error:", err)
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, e.ErrAccountNotExists
// 		}
// 		return nil, err
// 	}

// 	var transactions []entities.Account
// 	for rows.Next() {
// 		transaction := entities.Account{}
// 		err := rows.Scan(
// 			&transaction.ID,
// 			&transaction.CurrencyCode,
// 			&transaction.Balance,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		transactions = append(transactions, transaction)
// 	}

// 	return transactions, nil
// }

// // Create
// func (r *transactionRepo) Create(ctx context.Context, req entities.Transaction) (*entities.TransactionOut, error) {
// 	var (
// 		accountFromID sql.NullString
// 		accountToID   sql.NullString
// 	)
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer func() {
// 		if err != nil {
// 			fmt.Println("error:", err)
// 			tx.Rollback()
// 		} else {
// 			err = tx.Commit()
// 		}
// 	}()

// 	switch req.TypeID {
// 	case constants.TopupTransactionID:
// 		accountFromID = sql.NullString{Valid: false}
// 		accountToID.Scan(req.AccountToID)
// 	case constants.WithdrawTransactionID:
// 		accountToID = sql.NullString{Valid: false}
// 		accountFromID.Scan(req.AccountFromID)
// 	case constants.TransferTransactionID:
// 		accountToID.String = req.AccountToID
// 		accountFromID.String = req.AccountFromID
// 	}

// 	row := tx.QueryRowContext(ctx, `
// 	    INSERT INTO transactions (
// 			id,
// 			type_id,
// 			account_from_id,
// 			account_to_id,
// 			amount,
// 			created_at
// 		 )
// 		 VALUES ($1,
// 			 $2,
// 			 $3,
// 			 $4,
// 			 $5,
// 			 CURRENT_TIMESTAMP)
// 		 RETURNING
// 		 	id,
// 			(SELECT tt.name
// 				FROM transaction_types tt
// 				WHERE tt.id=type_id) AS typeName,
// 			account_from_id,
// 			account_to_id,
// 			amount,
// 			created_at
// 	`, req.ID,
// 		req.TypeID,
// 		accountFromID,
// 		accountToID,
// 		req.Amount)

// 	tr := entities.TransactionOut{}
// 	err = row.Scan(
// 		&tr.ID,
// 		&tr.TypeName,
// 		&accountFromID,
// 		&accountToID,
// 		&tr.Amount,
// 		&tr.CreatedAt,
// 	)
// 	if err != nil {
// 		pgErr, isPGErr := err.(*pq.Error)
// 		if isPGErr {
// 			if pgErr.Code == constants.PGForeignKeyViolationCode {
// 				return nil, e.ErrAccountNotExists
// 			}
// 		}
// 		return nil, err
// 	}

// 	tr.AccountFromID = accountFromID.String
// 	tr.AccountToID = accountToID.String

// 	return &tr, nil
// }
