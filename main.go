package main

import "github.com/restuwahyu13/go-jwt-transform/helpers"

// Encrypt jwt token using caesar cipher cryptography from real jwt token into fake jwt token
func Encrypt(token string, rotate uint) (string, error) {
	err := helpers.NewValidator(token, rotate)
	return helpers.Rotate(token, int(rotate)), err
}

// Decrypt jwt token using caesar cipher cryptography from fake jwt token into real jwt token
func Decrypt(token string, rotate uint) (string, error) {
	err := helpers.NewValidator(token, rotate)
	res := helpers.Rotate(token, int(-rotate))
	return res, err
}
