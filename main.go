package transform

import (
	"github.com/restuwahyu13/go-jwt-transform/helpers"
)

// Encrypt jwt token using caesar cipher cryptography from real jwt token into fake jwt token
func Encrypt(token string, rotate uint, privatekey string) (string, error) {
	if err := helpers.NewValidator(token, rotate); err != nil {
		return "error", err
	}

	if err := helpers.NewCredentials(token, privatekey, "encrypt"); err != nil {
		return "error", err
	}

	res := helpers.Rotate(token, int(rotate), privatekey, "encrypt")
	return res, nil
}

// Decrypt jwt token using caesar cipher cryptography from fake jwt token into real jwt token
func Decrypt(token string, rotate uint, privatekey string) (string, error) {
	if err := helpers.NewValidator(token, rotate); err != nil {
		return "error", err
	}

	if err := helpers.NewCredentials(token, privatekey, "decrypt"); err != nil {
		return "error", err
	}

	res := helpers.Rotate(token, int(-rotate), privatekey, "decrypt")
	return res, nil
}
