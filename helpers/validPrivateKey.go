package helpers

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidPrivateKey(privateKey string) (bool, error) {
	if ok, _ := regexp.MatchString(`[^A-Za-z0-9]`, privateKey); ok == true {
		return false, errors.New(fmt.Sprintf("privatekey not valid %s", privateKey))
	}

	regex := regexp.MustCompile(`[A-Za-z]+[0-9]|[0-9][A-Za-z]`)
	match := regex.MatchString(privateKey)
	return match, nil
}
