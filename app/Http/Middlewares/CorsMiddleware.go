package middlewares

import (
	"net/http"
)

// CorsMiddleware is a middleware to handle Cross-Origin Resource Sharing (CORS) requests
func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set headers for allowed methods, headers, and origins
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Origin", "*")

        // Handle pre-flight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Call the next middleware/handler in the chain
        next.ServeHTTP(w, r)
    })
}
