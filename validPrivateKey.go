package transform

import (
	"fmt"
	"regexp"
	"strings"
)

func validPrivateKey(token, privateKey, typeRotate string) (bool, error) {
	newToken := ""

	if typeRotate == "encrypt" {
		newToken = token
	} else if typeRotate == "decrypt" {
		newToken = token
	}

	if ok, _ := regexp.MatchString(`[^A-Za-z0-9]`, privateKey); ok == true {
		return false, fmt.Errorf("privatekey not valid %s", privateKey)
	} else if strings.Contains(newToken, privateKey) {
		return false, fmt.Errorf("privatekey cannot use jwt token %s", privateKey)
	}

	regex := regexp.MustCompile(`[A-Za-z]+[0-9]|[0-9][A-Za-z]`)
	match := regex.MatchString(privateKey)
	return match, nil
}
