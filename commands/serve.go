package commands

import (
	"fmt"
	"net/http"
	"gonga/packages"
	"gonga/routes"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application on the Golang development server",
	Args:  cobra.ExactArgs(0),
	Run: func(_ *cobra.Command, _ []string) {
		// start server 
		router := packages.NewRouter()
		routes.RegisterApiRoutes(router)
		fmt.Println("Server started on port 8080")
		http.ListenAndServe(":8080", router)
	},
}
