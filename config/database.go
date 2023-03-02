package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Charset  string
}

func LoadDatabaseConfig() (*gorm.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	log.Println(username, password, host, port, database)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// return &DatabaseConfig{
//     Host:     "localhost",
//     Port:     "3306",
//     User:     "root",
//     Password: "",
//     Name:     "my_database",
//     Charset:  "utf8mb4",
// }
