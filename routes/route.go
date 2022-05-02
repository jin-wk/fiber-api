package routes

import (
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jin-wk/fiber-api/config"
	_ "github.com/jin-wk/fiber-api/docs"
	"github.com/jin-wk/fiber-api/handler"
	"github.com/jin-wk/fiber-api/middleware"
)

func InitRoute(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/docs/*", fiberSwagger.HandlerDefault)

	api := app.Group("/api", logger.New(middleware.LoggerConfig()))
	api.Post("/auth/register", handler.Register)
	api.Post("/auth/login", handler.Login)

	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Env("JWT_SECRET_KEY")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(401).JSON(&fiber.Map{
				"message": "Unauthorized",
				"data":    nil,
			})
		},
	}))
	api.Get("/auth/:id", handler.Info)
	api.Post("/boards", handler.CreateBoard)
}
