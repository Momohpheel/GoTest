package database

import (
	"github.com/go-interview/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {

	userSeeder(db)

}

func userSeeder(db *gorm.DB) {

	email := "philip@test.com"
	fullname := "Philip Test"

	user := models.User{Email: email}
	row := db.Where(user).First(&user)

	if row.RowsAffected == 0 {

		//create password hash
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("test1234"), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		userRecord := models.User{
			FullName:       fullname,
			Email:          email,
			AccountBalance: 0,
			Password:       string(passwordHash),
		}

		// pass pointer of data to Create
		result := db.Create(&userRecord)

		if result.RowsAffected == 0 {
			return
		}

	}
}
