package database

import (
	"fmt"
	"gonga/config"
	"gonga/utils"
	"os"

	"github.com/pterm/pterm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dbConfig := config.LoadDatabaseConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		pterm.Error.WithShowLineNumber(true).Println("Cannot connect to database.", err)
		utils.Marlboro.Println("Cannot connect to database.", err)
		/**
		 * Use panic when something horribly goes wrong, i.e. an error that should have been
		 * caught before going to production. This is why is prints the stack.
		 *
		 * Use os.Exit to terminate the application cleanly for users.
		 * https://stackoverflow.com/a/28473273/12478479
		 */
		os.Exit(0)
	}

	return db, nil
}
