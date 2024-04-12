package transform

import (
	"math"
	"regexp"
	"strings"
)

var (
	alphabet   string = "abcdefghijklmnopqrstuvwxyz"
	alphacount int    = len(alphabet)
	lca        string = strings.ToLower(alphabet)
	uca        string = strings.ToUpper(alphabet)
)

func rotation(token string, rotate int, privateKey, typeRotate string) string {
	tokenArr := strings.Split(token, ".")
	firstToken := tokenArr[0]
	middleToken := tokenArr[1]
	lastToken := tokenArr[2]
	mergeToken := ""

	if typeRotate == "encrypt" {
		mergeToken = firstToken + "." + middleToken + privateKey + "." + lastToken
	} else {
		mergeToken = firstToken + "." + middleToken + "." + lastToken
	}

	regex := regexp.MustCompile(`\s`)
	cleanupMergeToken := regex.ReplaceAllString(mergeToken, "")

	if typeRotate == "decrypt" {
		encryptPrivateKey := rotationPrivateKey(privateKey, int(math.Abs(float64(rotate))))

		if !strings.Contains(cleanupMergeToken, encryptPrivateKey) {
			panic("private key credentials not match")
		}

		regex := regexp.MustCompile(encryptPrivateKey)
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

	return string(text)
}

func rotationPrivateKey(privateKey string, rotate int) string {
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
