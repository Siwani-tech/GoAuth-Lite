package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil

}

func CheckPasshash(password string, hashedpass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(password))
}
