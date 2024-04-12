package transform

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func validator(token string, rotate uint) error {
	if token == "" {
		return fmt.Errorf("token required %s", token)
	} else if rotate <= 0 {
		return fmt.Errorf("rotate cannot zero value %d", rotate)
	} else if reflect.TypeOf(token).Kind() != reflect.String {
		return fmt.Errorf("token must be string format %s", token)
	} else if reflect.TypeOf(rotate).Kind() != reflect.Uint {
		return fmt.Errorf("rotate must be number format %d", rotate)
	}

	toArray := strings.Split(token, ".")
	if len(toArray) != 3 {
		return fmt.Errorf("token must be jwt format %s", token)
	}

	return nil
}

func credentials(token, privateKey, typeRorate string) error {
	if len(privateKey) <= 20 {
		return fmt.Errorf("privatekey length must be greater than 20 characters %d", len(privateKey))
	} else if ok, err := validPrivateKey(token, privateKey, typeRorate); !ok {
		return err
	}

	return nil
}

func validPrivateKey(token, privateKey, typeRotate string) (bool, error) {
	newToken := token

	if typeRotate == "encrypt" || typeRotate == "decrypt" {
		newToken = token
	}

	if ok, _ := regexp.MatchString(`[^A-Za-z0-9]`, privateKey); ok {
		return false, errors.New(fmt.Sprintf("privatekey not valid %s", privateKey))
	} else if strings.Contains(newToken, privateKey) {
		return false, errors.New(fmt.Sprintf("privatekey cannot use jwt token %s", privateKey))
	}

	regex := regexp.MustCompile(`[A-Za-z]+[0-9]|[0-9][A-Za-z]`)
	match := regex.MatchString(privateKey)

	return match, nil
}
