package errors

import (
	"net/http"
)

type status struct {
	code        int
	description string
}

// NewError returns a new error with status code and description
func NewError(code int, description string) error {
	return status{code: code, description: description}
}

// Error implements the error interface
func (s status) Error() string {
	return s.description
}

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

// ExtractStatusCode extracts the status code from the error
func ExtractStatusCode(err error) (int, bool) {
	switch err := err.(type) {
	case status:
		// status error
		return err.code, true
	default:
		// non-status error
		return http.StatusInternalServerError, false
	}
}
