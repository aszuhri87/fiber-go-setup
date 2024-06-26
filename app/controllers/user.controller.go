package controllers

import (
	"fiber-go/app/models"
	"fiber-go/app/repositories"
	"fiber-go/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ShowUser godoc
// @Summary      Show all users
// @Description  get string by ID
// @Security Bearer
// @Tags         users
// @Accept       json
// @Produce      json
// @Param search query string false "Search"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success      200  {object}  models.ListResponseOk
// @Failure      400  {object}  lib.ErrorResponse
// @Failure      404  {object}  lib.ErrorResponse
// @Failure      500  {object}  lib.ErrorResponse
// @Router       /api/user [get]
func GetUser(c *fiber.Ctx) error {
	user := []models.User{}
	result := []models.ResponseData{}
	raw := models.ResponseData{}

	c.BodyParser(&user)
	query_params := c.Queries()

	search := query_params["search"]
	page, _ := strconv.Atoi(query_params["page"])
	limit, _ := strconv.Atoi(query_params["limit"])

	data, err := repositories.GetUser(user, search, page, limit)
	for i := 0; i < len(data); i++ {
		raw = models.ResponseData{ID: data[i].ID, Name: data[i].Name, Username: data[i].Username}
		result = append(result, raw)
	}

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
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
	result := models.ResponseData{ID: data.ID, Name: data.Name, Username: data.Username}

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
}

// ShowUserFirst godoc
// @Summary      Show an User By ID
// @Description  get User by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  lib.ErrorResponse
// @Failure      404  {object}  lib.ErrorResponse
// @Failure      500  {object}  lib.ErrorResponse
// @Router       /user/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	id := uuid.MustParse(c.Params("id"))

	data, err := repositories.GetUserByID(user, id)
	result := models.ResponseData{ID: data.ID, Name: data.Name, Username: data.Username}

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
}

// UpdateUser godoc
// @Summary      Update an User
// @Description  Update
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Param data body models.Create true "The input todo struct"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  lib.ErrorResponse
// @Failure      404  {object}  lib.ErrorResponse
// @Failure      500  {object}  lib.ErrorResponse
// @Router       /user/{id} [put]
func PutUser(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	id := uuid.MustParse(c.Params("id"))

	data, err := repositories.UpdateUser(user, id)
	result := models.ResponseData{ID: data.ID, Name: data.Name, Username: data.Username}

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
}

// DeleteUser godoc
// @Summary      Delete an User
// @Description  Delete
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Success      200  {object}  models.StatusOk
// @Failure      400  {object}  lib.ErrorResponse
// @Failure      404  {object}  lib.ErrorResponse
// @Failure      500  {object}  lib.ErrorResponse
// @Router       /user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	id := uuid.MustParse(c.Params("id"))

	data, err := repositories.DeleteUser(user, id)
	result := models.ResponseData{ID: data.ID, Name: data.Name, Username: data.Username}

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
}
