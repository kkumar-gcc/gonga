package middlewares

import (
	"log"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        log.Printf("Incoming request %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

        // Call the next handler
        next.ServeHTTP(w, r)

        // Log outgoing response
        log.Printf("Outgoing response %s %s to %s", r.Method, r.URL.Path, r.RemoteAddr)
		
	})
}
