package database

import (
	"log"

	"github.com/go-interview/models"
)

func Migrate() {
	// Migrate the schema
	db := DB
	err := db.AutoMigrate(
		&models.User{},
		&models.Transaction{},
	)

	if err != nil {
		log.Fatal(err)
	}

	seed(db)
}
