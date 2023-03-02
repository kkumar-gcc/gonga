package migrations

import (
	"gonga/app/Models"
	"log"

	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) error {
    err := db.AutoMigrate(&Models.User{})
    if err != nil {
        return err
    }

    return nil
}

func AddEmailToUsersTable(db *gorm.DB) error {
    err := db.AutoMigrate(&Models.User{})
    if err != nil {
        return err
    }

    type EmailColumn struct {
        Email string `gorm:"unique"`
    }

    err = db.Model(&Models.User{}).AutoMigrate(&Models.User{Name: "krishan", Email: "1@2.com",})
	if err != nil {
		log.Fatal(err)
	}
    if err != nil {
        return err
    }

    return nil
}

// Add more migration functions as needed
