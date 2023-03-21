package entities

import (
	"errors"

	"github.com/khdoba/banking/pkg/utils"
)

// Income
type Income struct {
	TransactionImp
}

// Validate ...
func (i *Income) Validate() error {
	if !utils.IsValidUUID(i.GetAccountID()) {
		return errors.New("invalid AccountID: invalid uuid")
	}
	return nil
}
