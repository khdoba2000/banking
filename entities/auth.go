package entities

import (
	"errors"

	"github.com/khdoba/banking/pkg/utils"
)

// LoginReq ...
type LoginReq struct {
	Password    string
	PhoneNumber string
}

// Validate ...
func (req *LoginReq) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid PhoneNumber: must in format +99XXXXXXXXXX")
	}

	return utils.ValidatePassword(req.Password)
}

// LoginRes ...
type LoginRes struct {
	ID          string
	PhoneNumber string
	Name        string
	Tokens      Tokens
}

// SignupReq ...
type SignupReq struct {
	ID          string
	PhoneNumber string
	Name        string
	Password    string
}

// Validate ...
func (req *SignupReq) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid PhoneNumber: must in format +99XXXXXXXXXX")
	}

	return utils.ValidatePassword(req.Password)
}

// SignupRes ...
type SignupRes struct {
	ID     string
	Tokens Tokens
}

// Tokens jwt tokens for authorization
type Tokens struct {
	AccessToken  string
	RefreshToken string
}
