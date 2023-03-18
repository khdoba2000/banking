package utils

import (
	"errors"
	"net"
	"reflect"
	"regexp"
	"unicode"

	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsEmailValid ...
func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

var uzbPhoneRegex = regexp.MustCompile(`[+]{1}99{1}[0-9]{10}$`)

// IsPhoneValid validates phone number for Uzbekistan
func IsPhoneValid(p string) bool {
	return uzbPhoneRegex.MatchString(p)
}

// Validate for both email ad phone
func ValidatePhoneOrEmail(loginValue string) bool {

	valid := IsPhoneValid(loginValue)
	if valid {
		return true
	}

	return IsEmailValid(loginValue)
}

// ValidatePassword check if the password meets requirements of
// at least 8 characters, at least one alphabetic, and at least one number
func ValidatePassword(p string) error {
	if len(p) < 8 {
		return errors.New("invalid password: must be at least 8 characters")
	}
	if len(p) > 256 {
		return errors.New("invalid password: must be at most 256 characters")
	}
	hasAnyDigit := false
	hasAnyAlphabetic := false
	for _, c := range p {
		if unicode.IsDigit(c) {
			hasAnyDigit = true
		}

		if unicode.IsLetter(c) {
			hasAnyAlphabetic = true
		}
	}

	if !hasAnyDigit {
		return errors.New("invalid password: must have at least one digit")
	}

	if !hasAnyAlphabetic {
		return errors.New("invalid password: must have at least one alphabetic")
	}

	return nil
}

// IsNil checks if the interface i is nil (empty)
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
