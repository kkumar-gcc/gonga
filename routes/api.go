package routes

import (
    "gonga/packages"
	"gonga/app/Http/Controllers"
	"gonga/app/Http/Middlewares"
)

func RegisterApiRoutes(router *packages.MyRouter) {

	router.Use(middlewares.LogMiddleware)

	UserController := controllers.UserController{};

	router.Get("/users", UserController.Index,middlewares.AuthMiddleware)
	router.Get("/users/{id}", UserController.Show)
	router.Post("/users", UserController.Create)
	router.Put("/users/{id}", UserController.Update)
	router.Delete("/users/{id}", UserController.Delete)



	
}
