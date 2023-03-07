package commands

import (
	"fmt"
	"gonga/bootstrap"

	"github.com/spf13/cobra"
)

func AboutCmd(app *bootstrap.Application) *cobra.Command {
	return &cobra.Command{
		Use:   "about",
		Short: "Display basic information about your application",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("This is a sample application built with Go and Cobra")
		},
	}
}
