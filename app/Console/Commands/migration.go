package commands

import (
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
		Hidden: true,
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


