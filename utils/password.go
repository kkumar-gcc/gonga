package utils

import (
	"errors"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+,.?/:;{}[]|"
	allChars  = letters + numbers + symbols
	minLength = 8
)

// HashPassword returns the bcrypt hash of the given password
func HashPassword(password string) (string, error) {
	// Generate a hashed representation of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword verifies that the given password matches the provided hash
func VerifyPassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

// GeneratePassword generates a random password of the given length with a mix of letters, numbers, and symbols
func GeneratePassword(length int) (string, error) {
	if length < minLength {
		return "", errors.New("password length must be at least 8")
	}
	rand.Seed(time.Now().UnixNano())
	var password string
	for i := 0; i < length; i++ {
		password += string(allChars[rand.Intn(len(allChars))])
	}
	return password, nil
}

// IsStrongPassword checks if the given password is strong, meaning it is at least 8 characters long
// and contains at least one lowercase letter, one uppercase letter, one number, and one symbol
// func IsStrongPassword(password string) bool {
// 	hasMinLength := len(password) >= minLength
// 	hasLowercase := false
// 	hasUppercase := false
// 	hasNumber := false
// 	hasSymbol := false
// 	for _, char := range password {
// 		switch {
// 		case 'a' <= char && char <= 'z':
// 			hasLowercase = true
// 		case 'A' <= char && char <= 'Z':
// 			hasUppercase = true
// 		case '0' <= char && char <= '9':
// 			hasNumber = true
// 		case symbols[0] <= char && char <= symbols[len(symbols)-1]:
// 			hasSymbol = true
// 		}
// 	}
// 	return hasMinLength && hasLowercase && hasUppercase && hasNumber && hasSymbol
// }
