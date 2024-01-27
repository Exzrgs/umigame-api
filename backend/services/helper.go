package services

import (
	"golang.org/x/crypto/bcrypt"
)

func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}
