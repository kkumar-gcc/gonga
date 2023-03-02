package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)


var aboutCmd = &cobra.Command{
    Use:   "about",
    Short: "Display basic information about your application",
    Run: func(_ *cobra.Command, _ []string) {
        fmt.Println("This is a sample application built with Go and Cobra")
    },
}
