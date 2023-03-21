package entities

import (
	"errors"

	"github.com/khdoba/banking/pkg/utils"
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
