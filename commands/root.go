package commands

import (
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use: "command",
	Short: "A CLI tool to manage application.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(aboutCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(dbCmd)
	//make commands
	rootCmd.AddCommand(makeControllerCmd)
	rootCmd.AddCommand(makeModelCmd)
	rootCmd.AddCommand(makeMigrationCmd)
	rootCmd.AddCommand(makeMiddlewareCmd)
	rootCmd.AddCommand(serveCmd)
	
}
