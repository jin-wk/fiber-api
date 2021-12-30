package routes

import (
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jin-wk/fiber-api/controller"
	_ "github.com/jin-wk/fiber-api/docs"
	"github.com/jin-wk/fiber-api/middleware"
)

func InitRoute(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/docs/*", fiberSwagger.Handler)

	api := app.Group("/api")
	api.Use(logger.New(middleware.LoggerConfig()))

	api.Post("/auth", controller.Register)
	api.Get("/auth/:id", controller.Info)
}
