package service

import (
	"crypto/rand"
	"math/big"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// GenPass interface
//go:generate mockery --name=GenPass --filename genpass.go --outpkg genpass_mocks
type GenPass interface {
	GeneratePassword(length int) (string, error)
}

type genPass struct{}

// NewGenPass create new instance of GenPass
func NewGenPass() GenPass {
	return &genPass{}
}

// GeneratePassword generate random password
func (g *genPass) GeneratePassword(length int) (string, error) {
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		password[i] = charset[n.Int64()]

	}

	return string(password), nil
}
