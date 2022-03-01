package hasher

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(MD5_Encode(password)), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashedAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(MD5_Encode(password)))
}
