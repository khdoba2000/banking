package entities

import (
	"github.com/google/uuid"
)

// LoginReq ...
type LoginReq struct {
	Password    string
	PhoneNumber string
}

// LoginRes ...
type LoginRes struct {
	ID          uuid.UUID
	PhoneNumber string
	Name        string
}

// SignupReq ...
type SignupReq struct {
	PhoneNumber string
	Name        string
	Password    string
}

// SignupRes ...
type SignupRes struct {
	ID uuid.UUID
}
