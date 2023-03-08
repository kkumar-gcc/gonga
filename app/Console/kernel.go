package console

import (
	"gonga/app/Console/Commands"
	"gonga/bootstrap"
	"os"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)
var rootCmd = &cobra.Command{
	Use: "command",
	Short: "A CLI tool to manage application.",
}

func Run(app *bootstrap.Application,db *gorm.DB) error {
	rootCmd.AddCommand(commands.AboutCmd(app))
	rootCmd.AddCommand(commands.ListCmd(app))
	rootCmd.AddCommand(commands.EnvCmd(app))
	rootCmd.AddCommand(commands.DbCmd(app))
	//make commands
	rootCmd.AddCommand(commands.MakeControllerCmd(app))
	rootCmd.AddCommand(commands.MakeModelCmd(app))
	rootCmd.AddCommand(commands.MakeMigrationCmd(app))
	rootCmd.AddCommand(commands.MigrateCmd(app,db))
	rootCmd.AddCommand(commands.MakeMiddlewareCmd(app))
	rootCmd.AddCommand(commands.ServeCmd(app))

	return rootCmd.Execute()
}

func RegisterCommands(app *bootstrap.Application,db *gorm.DB){
	err2 := Run(app,db)
	if err2 != nil {
		os.Exit(1)
	}
}
