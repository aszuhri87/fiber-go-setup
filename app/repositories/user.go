package repositories

import (
	"fiber-go/app/models"
	"fiber-go/database"
	"fiber-go/lib"

	"github.com/google/uuid"
)

func GetUser(user []models.User, search string, page int, limit int) ([]models.User, error) {
	offset := (page - 1) * limit
	data := database.DB.Limit(limit).Offset(offset)

	if search != "" {
		data.Where("name LIKE ? OR username LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	data.Find(&user)
	if data.Error != nil {
		return user, data.Error
	}
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	hash, _ := lib.HashPassword(user.Password)
	input := models.User{ID: uuid.New(), Name: user.Name, Username: user.Username, Password: hash}

	err := database.DB.Create(&input)
	if err.Error != nil {
		return input, err.Error
	}
	return input, nil
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
