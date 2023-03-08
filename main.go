package main

import (
	"gonga/app/Console"
	"gonga/bootstrap"
)

func main() {
	bootstrap.LoadEnv()

	app := bootstrap.NewApplication()
    
	db,_:=app.ConnectDatabase()

	// Register the routes
	app.RegisterApiRoutes()
    
	// register commands
	console.RegisterCommands(app,db)
}
