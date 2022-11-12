package transform

import "testing"

func TestRotateToken(action *testing.T) {
	action.Run("Should be TestRotateToken - encrypt private key not valid", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := 15
		privateKey := "abc"

		res, err := rotateToken(token, rotate, privateKey, "encrypt")
		if err != nil {
			assert(t, res, nil)
		}
	})

	action.Run("Should be TestRotateToken - decrypt private key not valid", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := 15
		privateKey := "abc"

		res, err := rotateToken(token, rotate, privateKey, "decrypt")
		if err != nil {
			assert(t, res, nil)
		}
	})

	action.Run("Should be TestRotateToken - encrypt private key valid", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := 15
		privateKey := "abcdefghijklmnopqrstuvwxyz123456789"

		res, err := rotateToken(token, rotate, privateKey, "encrypt")
		if err != nil {
			assert(t, res, nil)
		}
	})

	action.Run("Should be TestRotateToken - decrypt private key valid", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := 15
		privateKey := "abcdefghijklmnopqrstuvwxyz123456789"

		res, err := rotateToken(token, rotate, privateKey, "decrypt")
		if err != nil {
			assert(t, res, nil)
		}
	})
}
