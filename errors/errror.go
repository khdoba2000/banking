package errors

import (
	"net/http"

	e "github.com/khdoba/banking/pkg/errors"
)

// var (
// 	ErrPasswordTooShort      = status{Code: http.StatusBadRequest, Description: "password is too short"}
// 	ErrPasswordTooLong       = status{Code: http.StatusBadRequest, Description: "password is too long"}
// 	ErrMustContainDigit      = status{Code: http.StatusBadRequest, Description: "password must contain at least 1 digit"}
// 	ErrMustContainAlphabetic = status{Code: http.StatusBadRequest, Description: "password must contain at least 1 alphabetic"}
// 	ErrEmailAddress          = status{Code: http.StatusBadRequest, Description: "invalid email address. valid e-mail can contain only latin letters, numbers, '@' and '.'"}
// 	ErrInvalidRequestBody    = status{Code: http.StatusBadRequest, Description: "invalid request body"}
// 	ErrIncorrectNameValue    = status{Code: http.StatusBadRequest, Description: "name must include minumum 3 charachters"}
// 	ErrAuthIncorrect         = status{Code: http.StatusForbidden, Description: "auth incorrect"}
// 	ErrAuthNotGiven          = status{Code: http.StatusBadRequest, Description: "auth not given"}
// )

var (
	ErrCustomerNotExists     = e.NewError(http.StatusNotFound, "customer not exists")
	ErrAccountNotExists      = e.NewError(http.StatusBadRequest, "account not exists")
	ErrCustomerAlreadyExists = e.NewError(http.StatusBadRequest, "customer with this phone number already exists")
	ErrAccountAlreadyExists  = e.NewError(http.StatusBadRequest, "account with this info already exists")
)
