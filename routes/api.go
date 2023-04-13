package routes

import (
	"gonga/app/Http/Controllers"
	"gonga/app/Http/Controllers/Auth"
	"gonga/app/Http/Middlewares"
	"gonga/packages"

	"gorm.io/gorm"
)

func RegisterApiRoutes(router *packages.MyRouter, db *gorm.DB) {

	router.Use(middlewares.LogMiddleware).StrictSlash(true)
	router.Use(middlewares.CorsMiddleware).StrictSlash(true)
	
	UserController := controllers.UserController{DB: db}
	LoginController := auth.LoginController{DB: db}
	RegisterController := auth.RegisterController{DB: db}

	router.Post("/login", LoginController.Create)
	router.Post("/register", RegisterController.Create)
	router.Get("/users", UserController.Index, middlewares.AuthMiddleware)
	router.Get("/users/{id}", UserController.Show, middlewares.AuthMiddleware)
	router.Post("/users", UserController.Create)
	router.Put("/users/{id}", UserController.Update)
	router.Delete("/users/{id}", UserController.Delete)

}
