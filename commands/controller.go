package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// make:controller command
var makeControllerCmd = &cobra.Command{
    Use:   "make:controller [controller name]",
    Short: "Create a new controller",
    Long:  "Create a new controller with the specified name.",
    Args:  cobra.MinimumNArgs(1),
    Run: func(_ *cobra.Command, args []string) {
        // Get the controller name from the command line arguments
        name := args[0]

        // If the name contains a '/' character, create the directory if it doesn't exist
        var dirPath string
        if strings.Contains(name, "/") {
            parts := strings.Split(name, "/")
            dirPath = filepath.Join(parts[:len(parts)-1]...)
            if err := os.MkdirAll(filepath.Join("app/Http/Controllers", dirPath), os.ModePerm); err != nil {
                fmt.Printf("Error creating directory: %s\n", err.Error())
                return
            }
            name = parts[len(parts)-1]
        }

        // Create the controller file
        file, err := os.Create(filepath.Join("app/Http/Controllers", dirPath, name+".go"))
        if err != nil {
            fmt.Printf("Error creating file: %s\n", err.Error())
            return
        }
        defer file.Close()

        // Write the default controller code to the file
        file.WriteString(fmt.Sprintf("package %s\n\nimport \"fmt\"\n\nfunc %s() {\n    fmt.Println(\"Hello from %s controller\")\n}",name, name, name))

        // Print success message
        fmt.Printf("Controller [app/Http/Controllers/%s.go] created successfully!\n",name)
		
    },
}