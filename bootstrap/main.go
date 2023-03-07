package bootstrap

import (
	"fmt"
	"gonga/packages"
	"gonga/routes"
	"net/http"
)

// Application represents the Golang application instance.
type Application struct {
	Router *packages.MyRouter
}

// NewApplication creates a new instance of the Golang application.
func NewApplication() *Application {
	app := &Application{
		Router: packages.NewRouter(),
	}
	return app
}

// RegisterRoutes registers the application's routes.
func (app *Application) RegisterRoutes() {
	routes.RegisterApiRoutes(app.Router)
}

// Run starts the Golang application.
func (app *Application) Run() error {
	fmt.Println("Server started on http://localhost:8000")
	return http.ListenAndServe(":8000", app.Router)
}
