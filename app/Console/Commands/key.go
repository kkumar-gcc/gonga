package commands

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gonga/bootstrap"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// make:controller command
func KeyGenerateCmd(app *bootstrap.Application) *cobra.Command {

	return &cobra.Command{
		Use:   "key:generate",
		Short: "Generate a new application key",
		Long:  "Generate a new application key",
		Args:  cobra.MinimumNArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			// Generate 32 bytes of random data
			key := make([]byte, 64)
			_, err := rand.Read(key)
			if err != nil {
				return 
			}

			// Encode the random data using base64
			keyString := base64.StdEncoding.EncodeToString(key)

			// Output the new key
			pterm.Info.Printf("Application key [%s] generated successfully.\n", keyString)
			// Append the new key to the .env file
			file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return 
			}
			defer file.Close()

			_, err = fmt.Fprintf(file, "\nAPP_KEY=%s\n", keyString)
			if err != nil {
				return 
			}
        },
	}
}
