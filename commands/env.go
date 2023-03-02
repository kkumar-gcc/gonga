package commands

import (
	"log"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print environment variables",
	Long:  "Print all environment variables and their values in a nice table format.",
	Run: func(_ *cobra.Command, _ []string) {
		// Get all the environment variables
	    err := godotenv.Load(".env")
		
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
		env,_:=godotenv.Read(".env");
		// Create a new table
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		// Set table headers
		t.AppendHeader(table.Row{"Variable", "Value"})

		// Add each environment variable to the table
		for key, e := range env {
			pair := table.Row{key, e}
			t.AppendRow(pair)
		}

		// Render the table
		t.Render()
	},
}
