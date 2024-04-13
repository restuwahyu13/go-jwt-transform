package jwttransform

import (
	"errors"
	"math"
	"regexp"
	"strings"
)

const (
	ENC = "encrypt"
	DEC = "decrypt"
)

var (
	alphabet       string = "abcdefghijklmnopqrstuvwxyz"
	alphacount     int    = len(alphabet)
	lca            string = strings.ToLower(alphabet)
	uca            string = strings.ToUpper(alphabet)
	firstSecretKey []byte = nil
)

func getRotatedSecretKey(secretKey string, rotate int) string {
	text := []byte(secretKey)

	for i, v := range text {
		if v >= 'a' && v <= 'z' {
			text[i] = lca[(int((26+(v-'a')))+rotate)%alphacount]
		} else if v >= 'A' && v <= 'Z' {
			text[i] = uca[(int((26+(v-'A')))+rotate)%alphacount]
		}
	}

	return string(text)
}

func rotation(secretKey, token string, rotate int, rotateType string) ([]byte, error) {
	tokenArr := strings.Split(token, ".")
	firstToken := tokenArr[0]
	middleToken := tokenArr[1]
	lastToken := tokenArr[2]
	mergeToken := ""

	if rotateType == ENC {

		if len(secretKey) > 0 {
			firstSecretKey = []byte(secretKey)
		}

		mergeToken = firstToken + "." + middleToken + secretKey + "." + lastToken
	} else {
		mergeToken = firstToken + "." + middleToken + "." + lastToken
	}

	regex := regexp.MustCompile(`\s`)
	cleanupMergeToken := regex.ReplaceAllString(mergeToken, "")

	if rotateType == DEC {
		encryptSecretKey := getRotatedSecretKey(secretKey, int(math.Abs(float64(rotate))))

		if !strings.Contains(cleanupMergeToken, encryptSecretKey) || !strings.Contains(secretKey, string(firstSecretKey)) {
			return nil, errors.New("secretKey not match")
		}

		regex := regexp.MustCompile(encryptSecretKey)
		validJwtToken := regex.ReplaceAllString(cleanupMergeToken, "")
		cleanupMergeToken = validJwtToken
	}

	rotate %= alphacount
	text := []byte(cleanupMergeToken)

	for i, v := range text {
		if v >= 'a' && v <= 'z' {
			text[i] = lca[(int((26+(v-'a')))+rotate)%alphacount]
		} else if v >= 'A' && v <= 'Z' {
			text[i] = uca[(int((26+(v-'A')))+rotate)%alphacount]
		}
	}

	return text, nil
}
