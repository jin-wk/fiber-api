package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jin-wk/fiber-api/database"
	"github.com/jin-wk/fiber-api/models"
	"github.com/jin-wk/fiber-api/utils"
)

// Register godoc
// @Summary     Register
// @Description Register User
// @Tags		Auth
// @Accept		application/json
// @Produce		application/json
// @Param		user body models.User true "user"
// @Success		201 {object} utils.Response
// @Failure		500 {object} utils.Response
// @Router		/api/auth [post]
func Register(c *fiber.Ctx) error {
	var user models.RegisterUser
	var result models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(utils.Error("Bad Request", err))
	}

	err := utils.Validate(&user)
	if err != nil {
		return c.Status(400).JSON(utils.Error("Bad Request", err))
	}

	if database.DB.Where("email = ?", user.Email).First(&user).Error == nil {
		return c.Status(409).JSON(utils.Error("Email already exists", nil))
	}

	if database.DB.Select("Email", "Password", "Name").Create(&user).Error != nil {
		return c.Status(500).JSON(utils.Error("Internal Server Error", nil))
	}

	database.DB.First(&result, user.ID)
	return c.Status(201).JSON(utils.Success(result))
}

// Info godoc
// @Summary     Info
// @Description Get Info User
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		id path int true "id"
// @Success		200 {object} utils.Response
// @Failure		404 {object} utils.Response
// @Router		/api/auth/{id} [get]
func Info(c *fiber.Ctx) error {
	var user models.User

	result := database.DB.First(&user, c.Params("id"))
	if result.Error != nil {
		return c.Status(404).JSON(utils.Error("Not Found", nil))
	}

	return c.Status(200).JSON(utils.Success(&user))
}
