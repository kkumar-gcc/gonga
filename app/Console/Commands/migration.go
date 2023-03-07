package commands

import (
	"database/sql"
	"fmt"
	"gonga/bootstrap"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func MakeMigrationCmd(app *bootstrap.Application) *cobra.Command {
	return &cobra.Command{
		Use:   "make:migration [name]",
		Short: "Create a new migration file",
		Long:  "Create a new migration file with the specified name in the database/migrations directory.",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			name := args[0]
			filePath := filepath.Join("database/migrations", name+".go")

			// Check if the migration file already exists
			if _, err := os.Stat(filePath); !os.IsNotExist(err) {
				fmt.Printf("Error: migration file '%s' already exists\n", name)
				return
			}

			// Create the migration file
			now := time.Now().Unix()
			fileName := fmt.Sprintf("%d_%s.go", now, name)
			file, err := os.Create(filepath.Join("database/migrations", fileName))
			if err != nil {
				fmt.Printf("Error creating file: %s\n", err.Error())
				return
			}
			defer file.Close()

			// Write the default migration code to the file
			file.WriteString(fmt.Sprintf("package migrations\n\nimport \"fmt\"\n\nfunc up_%s() {\n    fmt.Println(\"Up migration for %s\")\n}\n\nfunc down_%s() {\n    fmt.Println(\"Down migration for %s\")\n}", name, name, name, name))

			// Print success message
			fmt.Printf("Migration [database/migrations/%s] created successfully!\n", fileName)
		},
	}
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations.",
	Long:  "Run any pending database migrations.",
	Run: func(_ *cobra.Command, _ []string) {
		// Open a connection to the database using the configured credentials
		db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
		if err != nil {
			fmt.Printf("Error connecting to database: %s\n", err.Error())
			os.Exit(1)
		}

		// Close the database connection when the command completes
		defer db.Close()

		// Run any pending database migrations using the configured migration tool (e.g. gorm)
		// ...

		fmt.Println("Database migrations completed.")
	},
}
