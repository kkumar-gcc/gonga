package middlewares

import (
	"net/http"
	"time"

	"github.com/pterm/pterm"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        pterm.Println(pterm.Gray(time.Now().Format("2006-01-02 15:04:05")) +"  Incoming request "+pterm.Magenta(r.Method)+" "+r.URL.Path+" from "+r.RemoteAddr )
	
        // Call the next handler
        next.ServeHTTP(w, r)
	})
}
