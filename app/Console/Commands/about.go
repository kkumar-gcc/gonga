package commands

import (
	"fmt"
	"gonga/bootstrap"

	"github.com/spf13/cobra"
)

func AboutCmd(app *bootstrap.Application) *cobra.Command {
	return &cobra.Command{
		/**
		 * The name and signature of the console command.
		 *
		 * @var string
		 */
		Use: "about",

		/**
		 * The name and signature of the console command.
		 *
		 * @var string
		 */
		Short: "Display basic information about your application",

		/**
		 * The name and signature of the console command.
		 *
		 * @var string
		 */
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("This is a sample application built with Go and Cobra")
		},
	}
}
