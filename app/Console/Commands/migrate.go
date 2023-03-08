package commands

import (
	"fmt"
	"gonga/app/Models"
	"gonga/bootstrap"
	"log"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func MigrateCmd(app *bootstrap.Application, db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations.",
		Long:  "Run any pending database migrations.",
		Run: func(_ *cobra.Command, args []string) {
			// Open a connection to the database using the configured credentials
			
			// Register the migrations
			err := db.AutoMigrate(
				&Models.User{},
				&Models.Address{},
			)
			if err != nil {
				log.Fatalf("Error running migrations: %v", err)
			}

			// Run any pending database migrations using the configured migration tool (e.g. gorm)
			// ...

			fmt.Println("Database migrations completed.")
		},
	}
}
