package helpers

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyz"
var alphacount int = len(alphabet)
var lca string = strings.ToLower(alphabet)
var uca string = strings.ToUpper(alphabet)

func Rotate(token string, rotate int, privatekey, typeRotate string) string {
	tokenArr := strings.Split(token, ".")
	firstToken := tokenArr[0]
	middleToken := tokenArr[1]
	lastToken := tokenArr[2]
	mergeToken := ""

	if typeRotate == "encrypt" {
		mergeToken = firstToken + "." + middleToken + privatekey + "." + lastToken
	} else {
		mergeToken = firstToken + "." + middleToken + "." + lastToken
	}

	regex := regexp.MustCompile(`\s`)
	cleanupMergeToken := regex.ReplaceAllString(mergeToken, "")

	if typeRotate == "decrypt" {
		encryptPrivateKey := rotatePrivateKey(privatekey, int(math.Abs(float64(rotate))))

		if res := strings.Contains(cleanupMergeToken, encryptPrivateKey); !res {
			panic("private key credentials not match")
		}

		regex := regexp.MustCompile(fmt.Sprintf(`%s`, encryptPrivateKey))
		validJwtToken := regex.ReplaceAllString(string(cleanupMergeToken), "")
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

	return string(text)
}

func rotatePrivateKey(privateKey string, rotate int) string {
	text := []byte(privateKey)

	for i, v := range text {
		if v >= 'a' && v <= 'z' {
			text[i] = lca[(int((26+(v-'a')))+rotate)%alphacount]
		} else if v >= 'A' && v <= 'Z' {
			text[i] = uca[(int((26+(v-'A')))+rotate)%alphacount]
		}
	}

	return string(text)
}
