package transform

import "testing"

func TestValidPrivateKey(action *testing.T) {
	action.Run("Should be TestValidPrivateKey - encrypt private key not valid", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		privateKey := "abc"

		res, err := validPrivateKey(token, privateKey, "encrypt")
		if err != nil {
			assert(t, res, false)
		}
	})

	action.Run("Should be TestValidPrivateKey - decrypt private key not valid", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		privateKey := "abc"

		res, err := validPrivateKey(token, privateKey, "decrypt")
		if err != nil {
			assert(t, res, false)
		}
	})

	action.Run("Should be TestValidPrivateKey - encrypt private key use jwt token", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		privateKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

		res, err := validPrivateKey(token, privateKey, "encrypt")
		if err != nil {
			assert(t, res, false)
		}
	})

	action.Run("Should be TestValidPrivateKey - decrypt private key use jwt token", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		privateKey := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9"

		res, err := validPrivateKey(token, privateKey, "decrypt")
		if err != nil {
			assert(t, res, false)
		}
	})

	action.Run("Should be TestValidPrivateKey - encrypt private key not use jwt token", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		privateKey := "abcdefghijklmnopqrstuvwxyz123456789"

		res, err := validPrivateKey(token, privateKey, "encrypt")

		if err != nil {
			t.FailNow()
		}

		if err == nil {
			assert(t, res, true)
		}
	})

	action.Run("Should be TestValidPrivateKey - decrypt private key not use jwt token", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		privateKey := "abcdefghijklmnopqrstuvwxyz123456789"

		res, err := validPrivateKey(token, privateKey, "decrypt")

		if err != nil {
			t.FailNow()
		}

		if err == nil {
			assert(t, res, true)
		}
	})
}
