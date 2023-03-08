package bootstrap

import (
	"fmt"
	"gonga/app/Http/Middlewares"
	"gonga/database"
	"gonga/packages"
	"gonga/routes"
	"net/http"

	"gorm.io/gorm"
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
func (app *Application) RegisterApiRoutes() {

	app.Router.Use(middlewares.CorsMiddleware)
	app.Router.Use(middlewares.ThrottleMiddleware)

	routes.RegisterApiRoutes(app.Router)
}
// ConnectDatabase connects to database.
func (app *Application) ConnectDatabase() (*gorm.DB, error) {
	return database.Connect()
}
// Run starts the Golang application.
func (app *Application) Run() error {
	fmt.Println("Server started on [http://localhost:8000]")
	return http.ListenAndServe(":8000", app.Router)
}

