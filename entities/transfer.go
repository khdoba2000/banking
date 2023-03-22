package entities

import (
	"errors"

	"github.com/khdoba2000/banking/constants"
	"github.com/khdoba2000/banking/pkg/utils"
)

// Transfer
type Transfer struct {
	TransactionImp
	AccountToID string
}

// Validate ...
func (t *Transfer) Validate() error {
	if !utils.IsValidUUID(t.AccountID) {
		return errors.New("invalid AccountID: invalid uuid")
	}

	if !utils.IsValidUUID(t.AccountToID) {
		return errors.New("invalid AccountID: invalid uuid")
	}

	if t.AccountID == t.AccountToID {
		return errors.New("cannot do transfer to the same account")
	}
	return nil
}

// GetTypeID ...
func (t *Transfer) GetTypeID() int {
	return constants.TransferTransactionID
}

// IsOut ...
func (t *Transfer) IsOut() bool {
	return true
}
