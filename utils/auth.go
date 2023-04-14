package utils

import (
	"gonga/app/Models"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/schema"
	"gorm.io/gorm"
)

func IsAuthenticate(r *http.Request) bool {
	// Check if user is authenticated
	tokenString := r.Header.Get("token")

	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide a signing key
		return []byte(Env("APP_KEY", "my-secret-key")), nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}

func Authenticate(username, password string, db *gorm.DB) (int, error) {
	// Get user by username from database
	var user Models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}
	
	// Compare passwords
	if err := VerifyPassword(user.Password, password); err != nil {
		log.Println(err)
		return 0, err
	}

	return int(user.ID), nil
}

func GenerateToken(userID int) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Get the secret key from an environment variable
	secretKey := []byte(Env("APP_KEY", "my-secret-key"))

	// Generate the signed token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil // Return a jwt token
}

func DecodeRequestBody(r *http.Request, v interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	err = decoder.Decode(v, r.PostForm)
	if err != nil {
		return err
	}

	return nil
}
