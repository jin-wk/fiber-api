package utils

import "github.com/gofiber/fiber/v2"

type Res struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(Res{message, data})
}
