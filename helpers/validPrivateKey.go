package helpers

import (
	"regexp"
)

func ValidPrivateKey(privateKey string) bool {
	regex := regexp.MustCompile(`[A-Za-z]+[0-9]|[0-9][A-Za-z]`)
	match := regex.MatchString(privateKey)
	return match
}
