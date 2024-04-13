package jwttransform

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func validator(secretKey, token string, rotate int, rotateType string) error {
	tokenType := "plainText"
	if rotateType != ENC {
		tokenType = "cipherText"
	}

	if len(token) <= 0 {
		return errors.New(fmt.Sprintf("%s not to be a empty", tokenType))
	} else if reflect.TypeOf(token).Kind() != reflect.String {
		return errors.New(fmt.Sprintf("%s must be a string format", tokenType))
	} else if len(strings.Split(token, ".")) != 3 {
		return errors.New(fmt.Sprintf("%s must be a jwt format", tokenType))
	} else if rotate <= 0 {
		return errors.New("rotate not to be a empty")
	} else if reflect.TypeOf(rotate).Kind() != reflect.Int {
		return errors.New("rotate must be a number format")
	} else if len(secretKey) <= 20 {
		return errors.New("secretKey length must be a greater than 20 characters")
	}

	return nil
}

func validSecretKey(secretKey, token string, rotate int, rotateType string) error {
	tokenType := "plainText"
	if rotateType != ENC {
		tokenType = "cipherText"
	}

	if err := validator(secretKey, token, rotate, rotateType); err != nil {
		return err
	}

	newToken := token
	if rotateType == ENC || rotateType == DEC {
		newToken = token
	}

	validSecretKey, err := regexp.MatchString(`[^A-Za-z0-9]`, secretKey)
	if err != nil {
		return err
	} else if validSecretKey {
		return errors.New("secretKey invalid format")
	} else if strings.Contains(newToken, secretKey) {
		return errors.New(fmt.Sprintf("secretKey cannot use %s", tokenType))
	}

	regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !regex.MatchString(secretKey) {
		return errors.New("secretKey invalid format")
	}

	return nil
}
