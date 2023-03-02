package middlewares

import (
	"log"
	"net/http"
)

func DefaultMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        log.Println("hello there routes")
        next.ServeHTTP(w, r)
    })
}