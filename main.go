package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jin-wk/go-test/config"
	"github.com/jin-wk/go-test/routes"
)

// @title         Fiber-API
// @version       0.0.1
// @description   Fiber Web API
// @contact.name  jin-wk
// @contact.url   https://github.com/jin-wk
// @contact.email note@kakao.com
// @host          localhost:5000
// @BasePath      /
func main() {
	if err := config.InitDatabase(); err != nil {
		log.Panic("Can't Connect Database: ", err.Error())
	}
	app := fiber.New(fiber.Config{})
	routes.InitRoute(app)
	log.Fatal(app.Listen(":5000"))
}
