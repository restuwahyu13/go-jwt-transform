package helpers

import "strings"

var alphabet string = "abcdefghijklmnopqrstuvwxyz"
var alphacount int = len(alphabet)
var lca string = strings.ToLower(alphabet)
var uca string = strings.ToUpper(alphabet)

func Rotate(token string, rotate int) string {
	rotate %= alphacount
	text := []byte(token)

	for i, v := range text {
		if v >= 'a' && v <= 'z' {
			text[i] = lca[(int((26+(v-'a')))+rotate)%alphacount]
		} else if v >= 'A' && v <= 'Z' {
			text[i] = uca[(int((26+(v-'A')))+rotate)%alphacount]
		}
	}

	return string(text)
}
