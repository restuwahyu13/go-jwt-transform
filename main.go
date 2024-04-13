package jwttransform

// Transform jwt token using caesar cipher cryptography from real jwt token into fake jwt token
func Transform(secretKey string, plainText string, rotate int) ([]byte, error) {
	if err := validSecretKey(secretKey, plainText, rotate, ENC); err != nil {
		return nil, err
	}
	return rotation(secretKey, plainText, rotate, ENC)
}

// Utransform jwt token using caesar cipher cryptography from fake jwt token into real jwt token
func Untransform(secretKey string, token string, rotate int) ([]byte, error) {
	if err := validSecretKey(secretKey, token, rotate, DEC); err != nil {
		return nil, err
	}
	return rotation(secretKey, token, -rotate, DEC)
}
