package main

import (
	"fmt"

	"github.com/restuwahyu13/go-jwt-transform/helpers"
)

// Encrypt jwt token using caesar cipher cryptography from real jwt token into fake jwt token
func Encrypt(token string, rotate uint, privatekey string) (string, error) {
	err := helpers.NewValidator(token, rotate)
	return helpers.Rotate(token, int(rotate), privatekey, "encrypt"), err
}

// Decrypt jwt token using caesar cipher cryptography from fake jwt token into real jwt token
func Decrypt(token string, rotate uint, privatekey string) (string, error) {
	err := helpers.NewValidator(token, rotate)
	res := helpers.Rotate(token, int(-rotate), privatekey, "decrypt")
	return res, err
}

func main() {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	privateKey := "27f06382c0645033294b7bc10250dd1ed9cc6bc5"

	encrypt, _ := Encrypt(accessToken, 15, privateKey)
	fmt.Println(encrypt)

	decrypt, _ := Decrypt(encrypt, 15, privateKey)
	fmt.Println(decrypt)
}
