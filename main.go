package transform

// Encrypt jwt token using caesar cipher cryptography from real jwt token into fake jwt token
func Encrypt(token string, rotate uint, privatekey string) (string, error) {
	if err := validator(token, rotate); err != nil {
		return "error", err
	}

	if err := credentials(token, privatekey, "encrypt"); err != nil {
		return "error", err
	}

	res, err := rotateToken(token, int(rotate), privatekey, "encrypt")
	if err != nil {
		return "error", err
	}

	return res.(string), nil
}

// Decrypt jwt token using caesar cipher cryptography from fake jwt token into real jwt token
func Decrypt(token string, rotate uint, privatekey string) (string, error) {
	if err := validator(token, rotate); err != nil {
		return "error", err
	}

	if err := credentials(token, privatekey, "decrypt"); err != nil {
		return "error", err
	}

	res, err := rotateToken(token, int(-rotate), privatekey, "decrypt")
	if err != nil {
		return "error", err
	}

	return res.(string), nil
}
