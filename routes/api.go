package routes

import (
	"gonga/app/Http/Controllers"
	"gonga/app/Http/Middlewares"
	"gonga/packages"

	"gorm.io/gorm"
)

func RegisterApiRoutes(router *packages.MyRouter,db *gorm.DB) {

	router.Use(middlewares.LogMiddleware)

	UserController := controllers.UserController{DB:db};

	router.Get("/users", UserController.Index)
	router.Get("/users/{id}", UserController.Show)
	router.Post("/users", UserController.Create)
	router.Put("/users/{id}", UserController.Update)
	router.Delete("/users/{id}", UserController.Delete)
	
}
