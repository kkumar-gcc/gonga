package routes

import (
	controllers "gonga/app/Http/Controllers"
	auth "gonga/app/Http/Controllers/Auth"
	middlewares "gonga/app/Http/Middlewares"
	"gonga/packages"

	"gorm.io/gorm"
)

// @title Gonga Api
// @description A social media api.
// @version 1.0
// @contact.name API Support
// @host localhost:8000
// @BasePath /
func RegisterApiRoutes(router *packages.MyRouter, db *gorm.DB) {

	UserController := controllers.UserController{DB: db}
	LoginController := auth.LoginController{DB: db}
	RegisterController := auth.RegisterController{DB: db}

	// Login API endpoint handlers
	router.Post("/login", LoginController.Create)

	// Register API endpoint handlers
	router.Post("/register", RegisterController.Create)

	// User API endpoint handlers
	router.Get("/users", UserController.Index, middlewares.AuthMiddleware)
	router.Get("/users/{id}", UserController.Show, middlewares.AuthMiddleware)
	router.Post("/users", UserController.Create)
	router.Put("/users/{id}", UserController.Update)
	router.Delete("/users/{id}", UserController.Delete)

}
