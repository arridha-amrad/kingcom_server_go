package authservice

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *authService) HashPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error on hashing password, %v", err)
	}
	return string(hashedPassword), nil
}

func (s *authService) VerifyPassword(hashedPassword string, plainPassword string) error {
	return bcrypt.
		CompareHashAndPassword(
			[]byte(hashedPassword),
			[]byte(plainPassword),
		)
}
