package Stubs

import (
	"fmt"
)

func GetMiddlewareStub(middlewareName string) string {
    template := `package middlewares

import (
    "fmt"
    "net/http"
)

func %s(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Hello from %s.")
        next.ServeHTTP(w, r)
    })
}
`
    return fmt.Sprintf(template, middlewareName, middlewareName)
	
}
