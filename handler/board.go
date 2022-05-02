package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jin-wk/fiber-api/database"
	"github.com/jin-wk/fiber-api/models"
	"github.com/jin-wk/fiber-api/utils"
)

// Register godoc
// @Summary     Create
// @Description Create Board
// @Tags		Board
// @Accept		json
// @Produce		json
// @Param		user body models.CreateBoard true "board"
// @Security    Authorization
// @Router		/api/boards [post]
func CreateBoard(c *fiber.Ctx) error {
	var create models.CreateBoard

	if err := c.BodyParser(&create); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Bad Request",
			"data":    err,
		})
	}

	validate := utils.Validate(&create)
	if validate != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Bad Request",
			"data":    validate,
		})
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	create.UserId = int(claims["id"].(float64))
	if database.DB.Table("boards").Select("user_id", "title", "content").Create(&create).Error != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"message": "Created",
		"data":    &create,
	})
}
