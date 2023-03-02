package commands

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List commands",
    Run: func(cmd *cobra.Command, _ []string) {
        fmt.Println("Available commands:")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		// Set table headers
		t.AppendHeader(table.Row{"Command", "Description"})

		// Add each environment variable to the table
		for _, e := range cmd.Root().Commands() {
			pair := table.Row{e.Name(), e.Short}
			t.AppendRow(pair)
		}
		
		// Render the table
		t.Render()
    },
}
