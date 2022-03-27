package routes

import (
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jin-wk/fiber-api/config"
	"github.com/jin-wk/fiber-api/controllers"
	_ "github.com/jin-wk/fiber-api/docs"
	"github.com/jin-wk/fiber-api/middleware"
	"github.com/jin-wk/fiber-api/utils"
)

func InitRoute(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/docs/*", fiberSwagger.Handler)

	api := app.Group("/api", logger.New(middleware.LoggerConfig()))
	api.Post("/auth/register", controllers.Register)
	api.Post("/auth/login", controllers.Login)

	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Env("JWT_SECRET_KEY")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Response(c, 401, "Unauthorized", nil)
		},
	}))
	api.Get("/auth/:id", controllers.Info)
}