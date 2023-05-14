package packages

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Route interface {
    Handle(method string, path string, handler HandlerFunc)
    Get(path string, handler HandlerFunc)
    Post(path string, handler HandlerFunc)
    Put(path string, handler HandlerFunc)
    Delete(path string, handler HandlerFunc)
}

// MyRouter is a wrapper around Gorilla mux.Router which provides an easier and cleaner way to define routes and middleware.
type MyRouter struct {
    *mux.Router
}

// NewRouter returns a new instance of MyRouter.
func NewRouter() *MyRouter {
    return &MyRouter{
        Router: mux.NewRouter(),
    }
}


// Handle adds a new route with a specified HTTP method to the router.
func (r *MyRouter) Handle(method string, path string, handler http.HandlerFunc) {
    r.Router.HandleFunc(path, handler).Methods(method)

}

// Use adds a middleware to the router.
func (r *MyRouter) Use(middleware func(http.Handler) http.Handler) *MyRouter {
    r.Router.Use(middleware)
    return r
}

// Subrouter returns a new instance of MyRouter with the same configuration as the parent route.
func (r *MyRouter) Subrouter(path string) *MyRouter {
	// initialize a subrouter with a copy of the parent route's configuration
	r.PathPrefix(path).Subrouter()
	return r
}


// Get adds a new route with the GET HTTP method and the specified path and handler function
// 
// Example usage: 
//   router.Get("/users", getUsersHandler)
func (r *MyRouter) Get(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter {
    r.Handle("GET", path, applyMiddleware(handler, middleware...))
    return r
}


// Post adds a new route with the POST HTTP method and the specified path and handler function
//
// Example usage: 
//   router.Post("/users", createUserHandler)
func (r *MyRouter) Post(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter{
    r.Handle("POST", path, applyMiddleware(handler, middleware...))
    return r
}


// Put adds a new route with the PUT HTTP method and the specified path and handler function
//
// Example usage:
//   router.Put("/users/{id}", updateUserHandler)
func (r *MyRouter) Put(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter{
    r.Handle("PUT", path, applyMiddleware(handler, middleware...))
    return r
}


// Delete adds a new route with the DELETE HTTP method and the specified path and handler function
//
// Example usage: 
//   router.Delete("/users/{id}", deleteUserHandler)
func (r *MyRouter) Delete(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter{
    r.Handle("DELETE", path, applyMiddleware(handler, middleware...))
    return r
}

// applyMiddleware applies a chain of middleware functions to an HTTP handler function
//
// If no middleware functions are provided, the original handler is returned.
//
// Parameters:
//   - h: the HTTP handler function to apply middleware to
//   - middleware: a variadic argument of middleware functions to apply to the handler, in order
//
// Returns:
//   - An HTTP handler function with the middleware applied.
//
// Example usage:
//   router.Get("/users", applyMiddleware(listUsersHandler, authMiddleware, loggingMiddleware))
func applyMiddleware(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
    if len(middleware) == 0 {
        return h
    }
    wrapped := h
    for i := len(middleware) - 1; i >= 0; i-- {
        wrapped = middleware[i](wrapped)
    }
    return wrapped
}
