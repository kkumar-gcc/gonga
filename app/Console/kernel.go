package console

import (
	"gonga/app/Console/Commands"
	"gonga/bootstrap"
	"os"

	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use: "command",
	Short: "A CLI tool to manage application.",
}

func Run(app *bootstrap.Application) error {
	rootCmd.AddCommand(commands.AboutCmd(app))
	rootCmd.AddCommand(commands.ListCmd(app))
	rootCmd.AddCommand(commands.EnvCmd(app))
	rootCmd.AddCommand(commands.DbCmd(app))
	//make commands
	rootCmd.AddCommand(commands.MakeControllerCmd(app))
	rootCmd.AddCommand(commands.MakeModelCmd(app))
	rootCmd.AddCommand(commands.MakeMigrationCmd(app))
	rootCmd.AddCommand(commands.MigrateCmd(app))
	rootCmd.AddCommand(commands.MakeMiddlewareCmd(app))
	rootCmd.AddCommand(commands.ServeCmd(app))

	return rootCmd.Execute()
}

func RegisterCommands(app *bootstrap.Application){
	err2 := Run(app)
	if err2 != nil {
		os.Exit(1)
	}
}
