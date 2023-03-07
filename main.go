package main

import (
	"gonga/app/console"
	"gonga/bootstrap"
)

func main() {
	
	bootstrap.LoadEnv()

	app:=bootstrap.NewApplication()

	// Register the routes
	app.RegisterRoutes()
	
	// register commands
	console.RegisterCommands(app)
}
