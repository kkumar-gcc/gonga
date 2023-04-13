package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

/**
 * Gets the value of an environment variable.
 *
 * @param  string  $key
 * @param  mixed  $default
 * @return mixed
 */
func Env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// JSONResponse sends a JSON response with the specified status code and data
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set status code
	w.WriteHeader(statusCode)

	// Marshal data to JSON format
	responseJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write response JSON to response writer
	w.Write(responseJSON)
}

// ErrorResponse represents the response format for validation errors.
type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// ValidateRequest validates the given request data against the specified validation rules
// using the Go Playground Validator library, and sends a formatted error response if validation fails.
func ValidateRequest(w http.ResponseWriter, rules interface{}) error {
	// Validate the form data using the validator library
	validate := validator.New()
	// Send a formatted error response if validation fails
	if err := validate.Struct(rules); err != nil {
		// Extract field names and error messages from the validation errors
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = ValidationMessage(e)
		}
		// Create the error response
		response := ErrorResponse{Errors: errors}
		// Send the error response
		JSONResponse(w, http.StatusBadRequest, response)
		return err
	}
	return nil
}
