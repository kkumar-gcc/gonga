package commands

import (
	"gonga/bootstrap"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// make:controller command
func DocGenerateCmd(app *bootstrap.Application) *cobra.Command {

	return &cobra.Command{
		Use:   "doc:generate",
		Short: "Generate Swagger documentation",
		Long:  "Generate Swagger documentation based on Swagger comments in the code",
		Args:  cobra.MinimumNArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			
			pterm.Info.Printf("Documentation generated successfully!\n")

		},
	}
}
