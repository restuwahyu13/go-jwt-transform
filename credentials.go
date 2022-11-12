package transform

import "fmt"

func credentials(token, privateKey, typeRorate string) error {
	if len(privateKey) <= 20 {
		return fmt.Errorf("privatekey length must be greater than 20 characters %d", len(privateKey))
	} else if ok, err := validPrivateKey(token, privateKey, typeRorate); !ok {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
