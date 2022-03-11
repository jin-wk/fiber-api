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
	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return utils.Response(c, 400, "Bad Request", err)
	}

	err := utils.Validate(user)
	if err != nil {
		return utils.Response(c, 400, "Bad Request", err)
	}

	if database.DB.Model(user).Where("email = ?", user.Email).First(&user).Error == nil {
		return utils.Response(c, 409, "Email Already Exists", nil)
	}

	if database.DB.Select("email", "password", "name").Create(&user).Error != nil {
		return utils.Response(c, 500, "Internal Server Error", nil)
	}

	return utils.Response(c, 201, "Created", models.ResponseUser{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
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
	var user models.ResponseUser

	result := database.DB.First(&user, c.Params("id"))
	if result.Error != nil {
		return utils.Response(c, 404, "Not Found", nil)
	}

	return utils.Response(c, 200, "OK", &user)
}
