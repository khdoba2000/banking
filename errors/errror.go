package errors

import (
	"net/http"

	e "github.com/khdoba2000/banking/pkg/errors"
)

var (
	ErrCustomerNotExists      = e.NewError(http.StatusNotFound, "customer not exists")
	ErrAccountNotExists       = e.NewError(http.StatusBadRequest, "account not exists")
	ErrCustomerAlreadyExists  = e.NewError(http.StatusBadRequest, "customer with this phone number already exists")
	ErrAccountAlreadyExists   = e.NewError(http.StatusBadRequest, "account with this info already exists")
	ErrInvalidTransactionType = e.NewError(http.StatusBadRequest, "invalid transaction type")
)
