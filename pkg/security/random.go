package security

import (
	"crypto/rand"
	"fmt"
)

// GenerateRandomBytes returns securely generated random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomCode returns a securely generated random string that consists of numbers that has length of 2*n
func GenerateRandomCode(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}

	code := ""
	for i := 0; i < n; i++ {
		code = code + fmt.Sprintf("%02d", b[i]%100)
	}

	return code, err
}
