package helpers

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func NewValidator(token string, rotate uint) error {
	return validator(token, rotate)
}

func validator(token string, rotate uint) error {
	typeofString := reflect.TypeOf(token)
	typeofNumber := reflect.TypeOf(rotate)

	if strings.Compare(token, "") != 1 {
		return fmt.Errorf("token required %s", token)
	} else if strings.Compare(fmt.Sprintf("%d", rotate), "0") != 1 {
		return fmt.Errorf("rotate cannot zero value %d", rotate)
	} else if typeofString != reflect.TypeOf("") {
		return fmt.Errorf("token must be string format %s", token)
	} else if typeofNumber != reflect.TypeOf(uint(time.Now().Year())) {
		return fmt.Errorf("rotate must be number format %d", rotate)
	}

	toArray := strings.Split(token, ".")
	if len(toArray) != 3 {
		return fmt.Errorf("token must be jwt format %s", token)
	}
	return nil
}
