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
//
// This function takes a plain text password as a string and generates its hashed representation
// using the bcrypt algorithm with the default cost. It returns the hashed password as a string.
// In case of an error, it returns an empty string and the corresponding error.
//
// Example usage:
//     hashedPassword, err := HashPassword("myPlainTextPassword")
//     if err != nil {
//         log.Fatal(err)
//     }
//
// Parameters:
//  - password (string): the plain text password to be hashed
//
// Returns:
//  - string: the bcrypt hash of the given password
//  - error: an error, if any, that occurred during the hashing process
func HashPassword(password string) (string, error) {
	// Generate a hashed representation of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


// VerifyPassword verifies that the given password matches the provided hash using the bcrypt algorithm.
//
// Example usage:
//	hash, _ := HashPassword("password")
//	err := VerifyPassword(hash, "password")
//	if err != nil {
//		fmt.Println("Invalid password!")
//	} else {
//		fmt.Println("Password is valid!")
//	}
//
// Parameters:
//   - hash (string): The hash of the password to be verified.
//   - password (string): The password to be verified.
//
// Returns:
//   - (error): Returns nil if the password matches the hash, otherwise returns an error.
func VerifyPassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}


// GeneratePassword generates a random password of the given length with a mix of letters,
// numbers, and symbols.
//
// If the length is less than 8(minLength), it returns an error.
//
// Example usage:
//     password, err := GeneratePassword(12)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(password)
//
// Parameters:
//   length (int): The length of the password to be generated.
//
// Returns:
//   string: The generated password.
//   error: An error, if any, that occurred during the password generation process.
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
