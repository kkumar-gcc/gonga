package main

import (
	console "gonga/app/Console"
	"gonga/bootstrap"
	"log"
)

var App *bootstrap.Application

func main() {
	bootstrap.LoadEnv()

	App = bootstrap.NewApplication()

	//connect to database
	err := App.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}

	// Register the routes
	App.RegisterApiRoutes()

	// register commands
	console.RegisterCommands(App)
}
