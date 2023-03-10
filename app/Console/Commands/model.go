package commands

import (
	"gonga/bootstrap"
	"gonga/packages/Stubs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func MakeModelCmd(app *bootstrap.Application) *cobra.Command {
	return &cobra.Command{
		Use:   "make:model [name]",
		Short: "Create a new model file.",
		Long:  "Create a new model with the specified name.",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			// Get the model name from the command argument
			name := args[0]

			// If the name contains a '/' character, create the directory if it doesn't exist
			var dirPath string
			if strings.Contains(name, "/") {
				parts := strings.Split(name, "/")
				dirPath = filepath.Join(parts[:len(parts)-1]...)
				if err := os.MkdirAll(filepath.Join("app/Models", dirPath), os.ModePerm); err != nil {
					pterm.Error.Printf("Error creating directory: %s\n", err.Error())
					return
				}
				name = parts[len(parts)-1]
			}

			// Create the controller file
			file, err := os.Create(filepath.Join("app/Models", dirPath, name+".go"))
			if err != nil {
				pterm.Error.Printf("Error creating file: %s\n", err.Error())
				return
			}
			defer file.Close()

			// Write the default controller code to the file
			file.WriteString(Stubs.GetModelStub(name))

			// Print success message
			pterm.Info.Printf("Model [app/Models/%s.go] created successfully!\n", name)
		},
	}
}
