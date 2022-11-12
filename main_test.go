package transform

import (
	"strings"
	"testing"
)

var privateKey string = "27f06382c0645033294b7bc10250dd1ed9cc6bc5"

func TestEncrypt(action *testing.T) {
	action.Run("Should be TestEncrypt - Token response encrypt match", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		output := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := uint(15)

		res, err := Encrypt(token, rotate, privateKey)

		if err != nil {
			t.FailNow()
		}

		assert(t, strings.Compare(res, output), 0)
	})
}

func TestDecrypt(action *testing.T) {
	action.Run("Should be TestEncrypt - Token response decrypt match", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		output := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(15)

		res, err := Decrypt(token, rotate, privateKey)

		if err != nil {
			t.FailNow()
		}

		assert(t, strings.Compare(res, output), 0)
	})
}
