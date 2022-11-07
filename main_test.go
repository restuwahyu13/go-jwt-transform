package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {

	t.Run("Error - Rotate must be not zero value", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(0)

		_, err := Encrypt(token, rotate)
		if err != nil {
			if err.Error() == fmt.Sprintf("rotate cannot zero value %d", rotate) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Token must be not empty", func(t *testing.T) {
		token := ""
		rotate := uint(0)

		_, err := Encrypt(token, rotate)
		if err != nil {
			if err.Error() == fmt.Sprintf("token required %s", token) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Token must be jwt format", func(t *testing.T) {
		token := "abcd"
		rotate := uint(15)

		_, err := Encrypt(token, rotate)
		if err != nil {
			if err.Error() == fmt.Sprintf("token must be jwt format %s", token) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Success - Token must be jwt format", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(15)

		_, err := Encrypt(token, rotate)
		if err == nil {
			t.Log("Success")
		}
	})

	t.Run("Success - Token response encrypt match", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(15)
		output := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"

		res, err := Encrypt(token, rotate)

		if err != nil {
			t.FailNow()
		}

		if strings.Compare(output, res) != 1 {
			t.Log("Success")
		} else {
			t.FailNow()
		}
	})
}

func TestDecrypt(t *testing.T) {

	t.Run("Error - Rotate must be not zero value", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(0)

		_, err := Decrypt(token, rotate)
		if err != nil {
			if err.Error() == fmt.Sprintf("rotate cannot zero value %d", rotate) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Token must be not empty", func(t *testing.T) {
		token := ""
		rotate := uint(0)

		_, err := Decrypt(token, rotate)
		if err != nil {
			if err.Error() == fmt.Sprintf("token required %s", token) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Token must be jwt format", func(t *testing.T) {
		token := "abcd"
		rotate := uint(15)

		_, err := Decrypt(token, rotate)
		if err != nil {
			if err.Error() == fmt.Sprintf("token must be jwt format %s", token) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Success - Token must be jwt format", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := uint(15)

		_, err := Decrypt(token, rotate)
		if err == nil {
			t.Log("Success")
		}
	})

	t.Run("Success - Token response decrypt match", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := uint(15)

		res, err := Decrypt(token, rotate)

		if err != nil {
			t.FailNow()
		}

		if strings.Compare(res, token) != 1 {
			t.Log("Success")
		} else {
			t.FailNow()
		}
	})
}
