package Utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPass := string(encryptPassword)
	if err != nil {
		return "", err
	}
	return hashPass, nil
}
func ComparePassword(dbPassword, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}
