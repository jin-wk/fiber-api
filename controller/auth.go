package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jin-wk/fiber-api/config"
	"github.com/jin-wk/fiber-api/model"
	"github.com/jin-wk/fiber-api/util"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary     Register
// @Description Register User
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		user body model.AddUser true "user"
// @Success		200 {object} model.Response
// @Failure		404 {object} model.Response
// @Failure		500 {object} model.Response
// @Router		/api/auth [post]
func Register(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(model.Error("Internal Server Error", nil))
	}

	err := util.Validate(user)
	if err != nil {
		return c.Status(400).JSON(model.Error("Bad Request", err))
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	create := config.DB.Create(user)
	if create.Error != nil {
		return c.Status(500).JSON(model.Error("Internal Server Error", nil))
	}

	return c.Status(200).JSON(model.Success(nil))
}

// Info godoc
// @Summary     Info
// @Description Get Info User
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		id path int true "id"
// @Success		200 {object} model.Response
// @Failure		404 {object} model.Response
// @Failure		500 {object} model.Response
// @Router		/api/auth/{id} [get]
func Info(c *fiber.Ctx) error {
	var user model.User
	result := config.DB.First(&user, c.Params("id"))
	if result.Error != nil {
		return c.Status(404).JSON(model.Error("Not Found", nil))
	}

	return c.Status(200).JSON(model.Success(user))
}
