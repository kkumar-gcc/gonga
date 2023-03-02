package main

import (
	"fmt"
	"gonga/commands"
	"gonga/config"
	"gonga/database/migrations"
	"os"
)

func main() {
	db, err := config.LoadDatabaseConfig()
    if err != nil {
        panic(err)
    }

    err = migrations.AddEmailToUsersTable(db)
    if err != nil {
        panic(err)
    }

    fmt.Println("Migration complete!")

	err2 := commands.Execute()
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
