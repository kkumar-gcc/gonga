package utils

import (
	"crypto/rand"
	"encoding/base64"
	"gonga/app/Models"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/schema"
	"gorm.io/gorm"
)

// IsAuthenticate checks if the user is authenticated by verifying the JWT token present in the request header.
//
// It expects the JWT token to be present in the "token" field of the request header. If it's not found, or if the token
// is invalid or expired, this function returns false. Otherwise, it returns true indicating that the user is authenticated.
//
// Example usage:
//     isAuthenticated := IsAuthenticate(r)
//     if isAuthenticated {
//         // perform authenticated operations
//     } else {
//         // redirect to login page or send unauthorized response
//     }
//
// Parameters:
//   r (*http.Request): The HTTP request object containing the token in its header.
//
// Returns:
//   bool: A boolean value indicating whether the user is authenticated or not.
func IsAuthenticate(r *http.Request) bool {
	// Check if user is authenticated
	tokenString := r.Header.Get("Authorization")

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


// Authenticate authenticates a user by checking the provided username and password against the database.
//
// It fetches the user with the given username from the database using the provided *gorm.DB object. If the user is found,
// it verifies the provided password against the user's hashed password using the VerifyPassword() function. If the password
// is correct, it returns the user ID as an integer. Otherwise, it returns an error indicating the reason for the failure.
//
// Example usage:
//     userID, err := Authenticate("john.doe", "my-password", db)
//     if err != nil {
//         log.Fatal(err)
//     }
//     // perform authenticated operations with userID
//
// Parameters:
//   username (string): The username of the user to be authenticated.
//   password (string): The password of the user to be authenticated.
//   db (*gorm.DB): The *gorm.DB object to be used for fetching the user from the database.
//
// Returns:
//   int: The user ID as an integer, if authentication is successful.
//   error: An error, if any, that occurred during the authentication process.
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


// UserExistsWithEmail checks if a user with the given email address exists in the database.
//
// It searches for the user with the provided email address in the database using the provided *gorm.DB object.
// If a user with the given email exists in the database, it returns true, otherwise it returns false.
//
// Example usage:
//     email := "john.doe@example.com"
//     exists := UserExistsWithEmail(db, email)
//     if exists {
//         fmt.Println("A user with the email", email, "already exists.")
//     } else {
//         fmt.Println("The email", email, "is available.")
//     }
//
// Parameters:
//   db (*gorm.DB): The *gorm.DB object to be used for fetching the user from the database.
//   email (string): The email address to be checked.
//
// Returns:
//   bool: True, if a user with the provided email exists in the database, otherwise false.
func UserExistsWithEmail(db *gorm.DB, email string) bool {
    var user Models.User
    result := db.Where("email = ?", email).First(&user)
    return result.Error == nil && result.RowsAffected > 0
}


// GenerateToken generates a new JWT token for the given user ID.
//
// This function generates a new JWT token with the given user ID as the "userID" claim and a default expiration time of 24 hours.
// The token is signed using the secret key retrieved from the "APP_KEY" environment variable. If the environment variable is not set,
// a default secret key of "my-secret-key" is used. The function returns the generated token as a string and an error, if any.
//
// Example usage:
//     userID := 1234
//     token, err := GenerateToken(userID)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println("Token:", token)
//
// Parameters:
//   userID (int): The ID of the user for whom the token is being generated.
//
// Returns:
//   string: The generated JWT token.
//   error: An error, if any, that occurred during the token generation process.
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


// DecodeRequestBody decodes the form data in the http request body and maps it to a struct
//
// This function takes an http request object and a struct to be decoded into, and returns an error
// if the decoding process fails. It first parses the request form data and then uses the
// github.com/gorilla/schema package to decode the form values into the given struct.
//
// Example usage:
//     var req MyRequestStruct
//     err := DecodeRequestBody(r, &req)
//     if err != nil {
//         log.Fatal(err)
//     }
//
// Parameters:
//  - r (*http.Request): the http request object containing the form data to be decoded
//  - v (interface{}): a pointer to the struct to be decoded into
//
// Returns:
//  - error: an error, if any, that occurred during the decoding process
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


// GenerateRandomString generates a random string of the given length using cryptographically secure random bytes
//
// This function takes an integer length and returns a random string of that length. The string is generated by
// generating random bytes using the rand package, and then encoding those bytes using base64 encoding. The resulting
// string is URL-safe, meaning it contains only URL-safe characters (i.e., it does not contain any '+' or '/' characters).
//
// Example usage:
//     randomString, err := GenerateRandomString(16)
//     if err != nil {
//         log.Fatal(err)
//     }
//
// Parameters:
//  - length (int): the length of the generated random string
//
// Returns:
//  - string: a random string of the given length
//  - error: an error, if any, that occurred during the generation process
func GenerateRandomString(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(bytes), nil
}

