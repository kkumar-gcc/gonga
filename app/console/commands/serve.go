package commands

import (
	"gonga/bootstrap"

	"github.com/spf13/cobra"
)

func ServeCmd(app *bootstrap.Application) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serve the application on the Golang development server",
		Args:  cobra.ExactArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			// start server
			app.Run()
		},
	}
}
