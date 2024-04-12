package transform

// Encrypt jwt token using caesar cipher cryptography from real jwt token into fake jwt token
func Encrypt(token string, rotate uint, secretKey string) (string, error) {
	if err := validator(token, rotate); err != nil {
		return "", err
	}

	if err := credentials(token, secretKey, "encrypt"); err != nil {
		return "", err
	}

	res := rotation(token, int(rotate), secretKey, "encrypt")
	return res, nil
}

// Decrypt jwt token using caesar cipher cryptography from fake jwt token into real jwt token
func Decrypt(token string, rotate uint, secretKey string) (string, error) {
	if err := validator(token, rotate); err != nil {
		return "", err
	}

	if err := credentials(token, secretKey, "decrypt"); err != nil {
		return "", err
	}

	res := rotation(token, int(-rotate), secretKey, "decrypt")
	return res, nil
}
