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
	DB *gorm.DB
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

	routes.RegisterApiRoutes(app.Router,app.DB)
}
// ConnectDatabase connects to database.
func (app *Application) ConnectDatabase() error {
    var err error
    app.DB, err = database.Connect()
    if err != nil {
        return err
    }
    return nil
}
// Run starts the Golang application.
func (app *Application) Run() error {
	fmt.Println("Server started on [http://localhost:8000]")
	return http.ListenAndServe(":8000", app.Router)
}

