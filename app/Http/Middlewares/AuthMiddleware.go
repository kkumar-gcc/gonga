package middlewares

import (
	"gonga/utils"
	"net/http"
)

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


