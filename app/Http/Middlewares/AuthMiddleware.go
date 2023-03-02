package middlewares

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if user is authenticated
		if !isAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// Call next middleware/handler
		next.ServeHTTP(w, r)
	})
}

func isAuthenticated(r *http.Request) bool {
	// Check if user is authenticated
	// ...
	return false
}
