package transform

import (
	"fmt"
	"strings"
	"testing"
)

var privateKey string = "27f06382c0645033294b7bc10250dd1ed9cc6bc5"

func TestEncrypt(t *testing.T) {

	t.Run("Error - Rotate must be not zero value", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(0)

		_, err := Encrypt(token, rotate, privateKey)
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
		rotate := uint(15)

		_, err := Encrypt(token, rotate, privateKey)
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

		_, err := Encrypt(token, rotate, privateKey)
		if err != nil {
			if err.Error() == fmt.Sprintf("token must be jwt format %s", token) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Private key must be greater than 20 character", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(15)
		privateKeyLength := len("abc123")

		_, err := Encrypt(token, rotate, "abc123")
		if err != nil {
			if err.Error() == fmt.Sprintf("privatekey length must be greater than 20 characters %d", privateKeyLength) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Private key not valid", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(15)
		newPrivateKey := fmt.Sprintf("===%s===", privateKey)

		_, err := Encrypt(token, rotate, newPrivateKey)
		if err != nil {
			if err.Error() == fmt.Sprintf("privatekey not valid %s", newPrivateKey) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Private key cannot same with jwt", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		rotate := uint(15)

		_, err := Encrypt(token, rotate, "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ")

		if err != nil {
			if err.Error() == fmt.Sprintf("privatekey cannot use jwt token %s", "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ") {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Success - Token response encrypt match", func(t *testing.T) {
		token := "eyJhbGciOiJSUzI1NiIsImI2NCI6dHJ1ZSwiY3R5IjoiSldUIiwia2lkIjoiSVNOZjJvMHphNyIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiNzI4NThiM2U3MyJdLCJleHAiOjE3MTI5NDgxNzAsImlhdCI6MTcxMjk0NDU3MCwiaXNzIjoiZmU3ZTFmMzNiNiIsImp0aSI6ImNmYTVkOWNmYmE3Y2VmNDFkMjIyMTU1NTVkNjNmNzEyZTcxNmRmYzkwMTY2ZjUxNjU5YWFmMzJjYTg3ZDM5ZDNkYWQ4ZjdjMDg1Yjc1MjdiODUyNjZhYWUwZmUwNzhkMzU5MTk0Nzk2MDFiZWRhNjQxOGU5NTQ3ZjVhZGNhMjRhIiwia2V5IjoiYmFiMTRmMzU2ZTM2NmIxYTVhOWRiNzBhOWJkMTVlZWRiZjY1M2Q1MGY1ZDJhMzM1NTYyMWE1MDBlYjc1MjU3ZDFhMzY1YjRmOTg4Yjg0MjkifQ.RiFO2BAeXGpFWWq3Y_PsDHM-J_KR4bda57fYToLmV1GbE9B8SBBeGfwTQShMaopFChOcYRRF5zgu-sCs1MvaAybTKr2f0yx00esZBt9JQPQc9R1tYf2GAlio8B9mwAwxIvqLQ7oztG8BbchBcLEuCUKb9vW8XGYG4KMOGkbHeXeszZnbjTVFO9hO0PnDZMcwXC1D84okP7YvuxEfDT8WPFtlLYYL7ryS_is0J5dZnUxnG8tfvKLN8lghytnD9feOHyGhh-uL9lguxUcEVauN98Rv4G8IOXzT59sKKH7iNgPaN02ltfE6NUjr6pPOl0wts22sp4eIjXPwPWHn5f2DJA"
		output := "tnYwqVrxDxYHJoX1CxXhXbX2CRX6sWY1OHlxN3G5XydxHasJXxlxp2azXydxHKCDOyYkBWewCnXhXcG5rRX6XzeMKRY9.tnYwsLFxDahxCoX4CIwxB2J3BnYsARYatWPxDyT3BIX5CSvmCoPhXbawsRX6BIrmByz0CSJ3BRlxpMCoXydxObJ3OIUbBoCxCxXhXbe0pHX6XbCbNIKzDLCbNbT3N2KbCSUzByXnBIJ1CIKzCyCbCoTnOIrmCbGbNozlBIN2OyJmCyJ5NLUbBoYyNIv3OSB5OSCzNLF4OysyBSv1Nyr1BysxDSJnCyOwNLJlObJlCowzBoJ5BIz0Coz2BSUxOLGwCyFmDVJ5CIF3OyKwOVCwByGwXxlxp2K5XydxNbUxBIGbBoJ2OIB2CbXmNIKwDLGxCoQwDLYzBIKaOLGxOyN1B2F1BVN1OSYwBoB1CINnBLT1BSQaNyr1ByJ3OSUwBoN1NyGbDIv4Nyv0ByzxuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.GxUD2QPtMVeULLf3N_EhSWB-Y_ZG4qsp57uNIdAbK1VqT9Q8HQQtVulIFHwBpdeURwDrNGGU5ovj-hRh1BkpPnqIZg2u0nm00thOQi9YFEFr9G1iNu2VPaxd8Q9blPlmXkfAF7doiV8QqrwQrATjRJZq9kL8MVNV4ZBDVzqWtMthoOcqyIKUD9wD0EcSOBrlMR1S84dzE7NkjmTuSI8LEUiaANNA7gnH_xh0Y5sOcJmcV8iukZAC8avwnicS9utDWnVww-jA9avjmJrTKpjC98Gk4V8XDMoI59hZZW7xCvEpC02aiuT6CJyg6eEDa0lih22he4tXyMElELWc5u2SYP"
		rotate := uint(15)

		res, err := Encrypt(token, rotate, privateKey)

		if err != nil {
			fmt.Print(err)
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
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF  27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := uint(0)

		_, err := Decrypt(token, rotate, privateKey)
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

		_, err := Decrypt(token, rotate, privateKey)
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

		_, err := Decrypt(token, rotate, privateKey)
		if err != nil {
			if err.Error() == fmt.Sprintf("token must be jwt format %s", token) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Private key must be greater than 20 character", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF  27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := uint(15)
		privateKeyLength := len("abc123")

		_, err := Encrypt(token, rotate, "abc123")
		if err != nil {
			if err.Error() == fmt.Sprintf("privatekey length must be greater than 20 characters %d", privateKeyLength) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Private key not valid", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF  27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := uint(15)
		newPrivateKey := fmt.Sprintf("===%s===", privateKey)

		_, err := Encrypt(token, rotate, newPrivateKey)

		if err != nil {
			if err.Error() == fmt.Sprintf("privatekey not valid %s", newPrivateKey) {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Error - Private key cannot same with jwt", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF  27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := uint(15)

		_, err := Encrypt(token, rotate, "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9")
		fmt.Printf("aaaa %s \n", err.Error())

		if err != nil {
			if err.Error() == fmt.Sprintf("privatekey cannot use jwt token %s", "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9") {
				t.Log("Success")
			} else {
				t.FailNow()
			}
		}
	})

	t.Run("Success - Token must be jwt format", func(t *testing.T) {
		token := "tnYwqVrxDxYHJoX1CxXhXbX2CRX6sWY1OHlxN3G5XydxHasJXxlxp2azXydxHKCDOyYkBWewCnXhXcG5rRX6XzeMKRY9.tnYwsLFxDahxCoX4CIwxB2J3BnYsARYatWPxDyT3BIX5CSvmCoPhXbawsRX6BIrmByz0CSJ3BRlxpMCoXydxObJ3OIUbBoCxCxXhXbe0pHX6XbCbNIKzDLCbNbT3N2KbCSUzByXnBIJ1CIKzCyCbCoTnOIrmCbGbNozlBIN2OyJmCyJ5NLUbBoYyNIv3OSB5OSCzNLF4OysyBSv1Nyr1BysxDSJnCyOwNLJlObJlCowzBoJ5BIz0Coz2BSUxOLGwCyFmDVJ5CIF3OyKwOVCwByGwXxlxp2K5XydxNbUxBIGbBoJ2OIB2CbXmNIKwDLGxCoQwDLYzBIKaOLGxOyN1B2F1BVN1OSYwBoB1CINnBLT1BSQaNyr1ByJ3OSUwBoN1NyGbDIv4Nyv0ByzxuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.GxUD2QPtMVeULLf3N_EhSWB-Y_ZG4qsp57uNIdAbK1VqT9Q8HQQtVulIFHwBpdeURwDrNGGU5ovj-hRh1BkpPnqIZg2u0nm00thOQi9YFEFr9G1iNu2VPaxd8Q9blPlmXkfAF7doiV8QqrwQrATjRJZq9kL8MVNV4ZBDVzqWtMthoOcqyIKUD9wD0EcSOBrlMR1S84dzE7NkjmTuSI8LEUiaANNA7gnH_xh0Y5sOcJmcV8iukZAC8avwnicS9utDWnVww-jA9avjmJrTKpjC98Gk4V8XDMoI59hZZW7xCvEpC02aiuT6CJyg6eEDa0lih22he4tXyMElELWc5u2SYP"
		rotate := uint(15)

		_, err := Decrypt(token, rotate, privateKey)
		if err == nil {
			t.Log("Success")
		}
	})

	t.Run("Success - Token response decrypt match", func(t *testing.T) {
		token := "tnYwqVrxDxYHJoX1CxXhXbX2CRX6sWY1OHlxN3G5XydxHasJXxlxp2azXydxHKCDOyYkBWewCnXhXcG5rRX6XzeMKRY9.tnYwsLFxDahxCoX4CIwxB2J3BnYsARYatWPxDyT3BIX5CSvmCoPhXbawsRX6BIrmByz0CSJ3BRlxpMCoXydxObJ3OIUbBoCxCxXhXbe0pHX6XbCbNIKzDLCbNbT3N2KbCSUzByXnBIJ1CIKzCyCbCoTnOIrmCbGbNozlBIN2OyJmCyJ5NLUbBoYyNIv3OSB5OSCzNLF4OysyBSv1Nyr1BysxDSJnCyOwNLJlObJlCowzBoJ5BIz0Coz2BSUxOLGwCyFmDVJ5CIF3OyKwOVCwByGwXxlxp2K5XydxNbUxBIGbBoJ2OIB2CbXmNIKwDLGxCoQwDLYzBIKaOLGxOyN1B2F1BVN1OSYwBoB1CINnBLT1BSQaNyr1ByJ3OSUwBoN1NyGbDIv4Nyv0ByzxuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.GxUD2QPtMVeULLf3N_EhSWB-Y_ZG4qsp57uNIdAbK1VqT9Q8HQQtVulIFHwBpdeURwDrNGGU5ovj-hRh1BkpPnqIZg2u0nm00thOQi9YFEFr9G1iNu2VPaxd8Q9blPlmXkfAF7doiV8QqrwQrATjRJZq9kL8MVNV4ZBDVzqWtMthoOcqyIKUD9wD0EcSOBrlMR1S84dzE7NkjmTuSI8LEUiaANNA7gnH_xh0Y5sOcJmcV8iukZAC8avwnicS9utDWnVww-jA9avjmJrTKpjC98Gk4V8XDMoI59hZZW7xCvEpC02aiuT6CJyg6eEDa0lih22he4tXyMElELWc5u2SYP"
		rotate := uint(15)

		res, err := Decrypt(token, rotate, privateKey)

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
