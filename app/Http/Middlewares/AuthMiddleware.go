package middlewares

import (
	"gonga/utils"
	"net/http"
)

// AuthMiddleware is a middleware function that checks if a user is authenticated.
// If the user is not authenticated, it returns an error response with status code 401.
// If the user is authenticated, it calls the next middleware/handler in the chain.
//
// Example usage:
//   router := packages.NewRouter()
//   router.Put("/api/users/{id}", UserController.Update, middlewares.AuthMiddleware)
//
// This will add the AuthMiddleware to the UserHandler function when the "/api/users" endpoint is accessed with the "GET" method.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if user is authenticated
		if ! utils.IsAuthenticate(r) {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}
		// Call next middleware/handler
		next.ServeHTTP(w, r)
	})
}


