package utils

import "github.com/gofiber/fiber/v2"

type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(Resp{message, data})
}
