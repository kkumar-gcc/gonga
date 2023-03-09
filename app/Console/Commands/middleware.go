package commands

import (
	"fmt"
	"gonga/bootstrap"
	"gonga/packages/Stubs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// make:controller command
func MakeMiddlewareCmd(app *bootstrap.Application) *cobra.Command {
	return &cobra.Command{
		Use:   "make:middleware [middleware name]",
		Short: "Create a new middleware",
		Long:  "Create a new middleware with the specified name.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {

			// Get the controller name from the command line arguments
			name := args[0]

			// If the name contains a '/' character, create the directory if it doesn't exist
			var dirPath string
			if strings.Contains(name, "/") {
				parts := strings.Split(name, "/")
				dirPath = filepath.Join(parts[:len(parts)-1]...)
				if err := os.MkdirAll(filepath.Join("app/Http/Middlewares", dirPath), os.ModePerm); err != nil {
					fmt.Printf("Error creating directory: %s\n", err.Error())
					return
				}
				name = parts[len(parts)-1]
			}

			// Create the controller file
			file, err := os.Create(filepath.Join("app/Http/Middlewares", dirPath, name+".go"))
			if err != nil {
				fmt.Printf("Error creating file: %s\n", err.Error())
				return
			}
			defer file.Close()

			// Write the default controller code to the file
			file.WriteString(Stubs.GetMiddlewareStub(name))

			// Print success message
			fmt.Printf("Middleware [app/Http/Middlewares/%s.go] created successfully!\n", name)

		},
	}
}
