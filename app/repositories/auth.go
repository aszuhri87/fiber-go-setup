package repositories

import (
	"fiber-go/app/models"
	"fiber-go/database"
	"fiber-go/lib"

	"github.com/google/uuid"
)

func RegisterUser(user models.User) (models.User, error) {
	hash, _ := lib.HashPassword(user.Password)

	input := models.User{ID: uuid.New(), Name: user.Name, Username: user.Username, Password: hash}

	err := database.DB.Create(&input)

	if err.Error != nil {
		return input, err.Error
	}
	return input, nil
}

func Login(user models.User) (models.User, error) {
	data := models.User{Username: user.Username}

	err := database.DB.First(&data)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}
