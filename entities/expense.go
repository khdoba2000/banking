package entities

import (
	"errors"

	"github.com/khdoba2000/banking/constants"
	"github.com/khdoba2000/banking/pkg/utils"
)

// Expense
type Expense struct {
	TransactionImp
}

// Validate ...
func (e *Expense) Validate() error {
	if !utils.IsValidUUID(e.GetAccountID()) {
		return errors.New("invalid AccountID: invalid uuid")
	}
	return nil
}

// GetTypeID ...
func (e *Expense) GetTypeID() int {
	return constants.ExpenseTransactionID
}

// IsOut ...
func (e *Expense) IsOut() bool {
	return true
}
