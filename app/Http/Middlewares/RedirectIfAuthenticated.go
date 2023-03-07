package middlewares

// import (
// 	"net/http"
// )

// func RedirectIfAuthenticated(next http.HandlerFunc) http.HandlerFunc {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         session, err := store.Get(r, "session-name")
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             return
//         }

//         if auth.IsAuthenticated(session) {
//             http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
//             return
//         }

//         next.ServeHTTP(w, r)
//     })
// }

