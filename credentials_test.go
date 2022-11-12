package transform

import (
	"fmt"
	"testing"
)

func TestCredentials(action *testing.T) {
	action.Run("Should be TestCredentials - encrypt private key length under 20 characters", func(t *testing.T) {
		token := ""
		privateKey := "abcd9"

		err := credentials(token, privateKey, "encrypt")

		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("privatekey length must be greater than 20 characters %d", len(privateKey)))
		}
	})

	action.Run("Should be TestCredentials - decrypt private key length under 20 characters", func(t *testing.T) {
		token := ""
		privateKey := "abcd9"

		err := credentials(token, privateKey, "decrypt")

		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("privatekey length must be greater than 20 characters %d", len(privateKey)))
		}
	})

	action.Run("Should be TestCredentials - encrypt private key valid", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		privateKey := "abcdefghijklmnopqrstuvwxyz123456789"

		err := credentials(token, privateKey, "encrypt")

		if err != nil {
			t.FailNow()
		}

		assert(t, nil, nil)
	})

	action.Run("Should be TestCredentials - decrypt private key valid", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		privateKey := "abcdefghijklmnopqrstuvwxyz123456789"

		err := credentials(token, privateKey, "decrypt")

		if err != nil {
			t.FailNow()
		}

		assert(t, nil, nil)
	})
}
