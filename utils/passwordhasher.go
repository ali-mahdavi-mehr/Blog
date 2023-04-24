package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err == nil {
		return true
	}
	return false
}
