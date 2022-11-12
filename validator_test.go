package transform

import (
	"fmt"
	"testing"
)

func TestValidator(action *testing.T) {
	action.Run("Should be TestValidator - token empty", func(t *testing.T) {
		token := ""
		rotate := 15

		err := validator(token, uint(rotate))
		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("token required %s", token))
		}
	})

	action.Run("Should be TestValidator - rotate not zero value", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := 0

		err := validator(token, uint(rotate))
		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("rotate cannot zero value %d", rotate))
		}
	})

	action.Run("Should be TestValidator - token not jwt format", func(t *testing.T) {
		token := "abcdefghijklmnopqrstuvwxyz123456789"
		rotate := 15

		err := validator(token, uint(rotate))
		if err != nil {
			assert(t, err.Error(), fmt.Sprintf("token must be jwt format %s", token))
		}
	})

	action.Run("Should be TestValidator - error not exist", func(t *testing.T) {
		token := "tnYwqVrxDxYXJoX1CxXhXcG5rRX6XzeMKRY9.tnYosLXxDxXmByB0CIN3DSzlXxlxqbUiOHX6XzekpV4vGV9aXxlxpLU0XydmCIT2ByB5BSXnuF27u06382r0645033294q7qr10250ss1ts9rr6qr5.HuaZmlGYHBtZZU2FI4uleBtYu36EDz6nYK_psFhhl5r"
		rotate := 15

		err := validator(token, uint(rotate))

		if err != nil {
			t.FailNow()
		}

		assert(t, err, nil)
	})
}
