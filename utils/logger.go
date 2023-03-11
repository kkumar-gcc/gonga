package utils

import (
	"log"
	"os"
	"path"

	"github.com/pterm/pterm"
)

/**
 * Define Custom Logger to implement log.Logger
 * The name is inspired by winstonjs/winston. A logger for just about everything.
 * Winston and Marlboro are both American brand of cigarettes.
 *
 * Marlboro Logger is designed to save important error logs.
 * Helpful for debugging in both development and production.
 */
var Marlboro *log.Logger

/**
 * Initialize Marlboro Logger instance
 */
func init() {
	/**
	 * Project root path
	 * This is where we run `go run main.go serve`
	 */
	root_path, _ := os.Getwd()

	/**
	 * Path to logs directory
	 * Default log folder is "storage/logs"
	 * At the moment there is no reason to provide custom log_folder for user.
	 */
	log_folder := path.Join(root_path, "storage/logs")

	log_file := path.Join(log_folder, "app.log")

	file, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		pterm.Error.Println("Failed to", err)
	}

	Marlboro = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}
