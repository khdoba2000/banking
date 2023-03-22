package entities

import (
	"errors"

	"github.com/khdoba2000/banking/pkg/utils"
)

// Customer ...
type Customer struct {
	ID          string
	PhoneNumber string
	Name        string
	Password    string
}

// Validate ...
func (req *Customer) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid PhoneNumber: must in format +99XXXXXXXXXX")
	}

	return utils.ValidatePassword(req.Password)
}
