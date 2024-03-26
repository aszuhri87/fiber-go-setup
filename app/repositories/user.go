package repositories

import (
	"fiber-go/app/models"
	"fiber-go/database"
	"fiber-go/lib"
	"fmt"

	"github.com/google/uuid"
)

func GetUser(user []models.User) ([]models.User, error) {
	data := database.DB.Find(&user)

	if data.Error != nil {
		return user, data.Error
	}

	return user, nil

}

func CreateUser(user models.User) (models.User, error) {
	hash, _ := lib.HashPassword(user.Password)
	fmt.Println(user)

	input := models.User{ID: uuid.New(), Name: user.Name, Username: user.Username, Password: hash}

	err := database.DB.Create(&input)

	if err.Error != nil {
		return input, err.Error
	}
	return input, nil

	// db := database.GetDB
	// sqlStatement := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3) RETURNING id`
	// err := db.QueryRow(sqlStatement, uuid.New(), user.Username, user.Password).Scan(&user.ID)
	// if err != nil {
	// 	return user, err
	// }
	// return user, nil

}

func GetUserByID(user models.User, id uuid.UUID) (models.User, error) {
	data := models.User{ID: id}

	err := database.DB.Find(&data)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}

func UpdateUser(user models.User, id uuid.UUID) (models.User, error) {
	data := models.User{ID: id}
	hash, _ := lib.HashPassword(user.Password)

	input := models.User{Name: user.Name, Username: user.Username, Password: hash}

	err := database.DB.Where("id = ?", id).Updates(&input)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}

func DeleteUser(user models.User, id uuid.UUID) (models.User, error) {
	data := models.User{ID: id}

	err := database.DB.Delete(&data, id)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}
