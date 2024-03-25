package controllers

import (
	"fiber-go/app/models"
	"fiber-go/app/repositories"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUser(c *fiber.Ctx) error {
	user := []models.User{}
	formated := []models.ResponseData{}
	raw := models.ResponseData{}

	c.BodyParser(&user)

	data, err := repositories.GetUser(user)
	for i := 0; i < len(data); i++ {
		raw = models.ResponseData{ID: data[i].ID, Name: data[i].Name, Username: data[i].Username}

	}
	formated = append(formated, raw)
	result := models.DataResponseOk{Data: formated, Message: "OK", Code: http.StatusOK}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(result)
}

// CreateUser godoc
// @Summary      Create an User
// @Description  create
// @Tags         users
// @Accept       json
// @Produce      json
// @Param data body models.Create true "The input todo struct"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  lib.ErrorResponse
// @Failure      404  {object}  lib.ErrorResponse
// @Failure      500  {object}  lib.ErrorResponse
// @Router       /user [post]
func PostUser(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	data, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Status(200)
	return c.JSON(data)
}

func GetUserByID(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	id := uuid.MustParse(c.Params("id"))

	data, err := repositories.GetUserByID(user, id)
	formated := models.ResponseData{ID: data.ID, Name: data.Name, Username: data.Username}
	result := models.DataResponseOk{Data: formated, Message: "OK", Code: http.StatusOK}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(result)
}

func PutUser(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	id := uuid.MustParse(c.Params("id"))

	data, err := repositories.UpdateUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	id := uuid.MustParse(c.Params("id"))

	data, err := repositories.DeleteUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}
