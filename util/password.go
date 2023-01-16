package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedstring, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedstring), err
}

func CheckHashedPassword(HashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))

	return err
}
