package main

import (
	console "gonga/app/Console"
	"gonga/bootstrap"
	"gonga/config"
	"log"
)

var App *bootstrap.Application

func main() {
	bootstrap.LoadEnv()
	mailConfig := config.LoadMailConfig()
	log.Println(mailConfig.Default,mailConfig.Mailers[mailConfig.Default].Port)
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
