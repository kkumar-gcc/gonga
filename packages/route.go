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
type MyRouter struct {
    *mux.Router
}

func NewRouter() *MyRouter {
    return &MyRouter{
        Router: mux.NewRouter(),
    }
}

func (r *MyRouter) Handle(method string, path string, handler http.HandlerFunc) {
    r.Router.HandleFunc(path, handler).Methods(method)

}
func (r *MyRouter) Use(middleware func(http.Handler) http.Handler) *MyRouter {
    r.Router.Use(middleware)
    return r
}
func (r *MyRouter) Subrouter(path string) *MyRouter {
	// initialize a subrouter with a copy of the parent route's configuration
	r.PathPrefix(path).Subrouter()
	return r
}

func (r *MyRouter) Get(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter {
    r.Handle("GET", path, applyMiddleware(handler, middleware...))
    return r
}

func (r *MyRouter) Post(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter{
    r.Handle("POST", path, applyMiddleware(handler, middleware...))
    return r
}

func (r *MyRouter) Put(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter{
    r.Handle("PUT", path, applyMiddleware(handler, middleware...))
    return r
}

func (r *MyRouter) Delete(path string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) *MyRouter{
    r.Handle("DELETE", path, applyMiddleware(handler, middleware...))
    return r
}


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
