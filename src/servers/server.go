package servers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"juanfer2/go-queries/src/modules/queries/routes"
)

type Server struct {
	Port string
}

func healt(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", "World")
	return c.SendString(msg) // => Hello john 👋!
}

func (s *Server) Start(port string) {
	app := fiber.New()
	app.Use(cors.New())
	s.Port = port

	app.Get("/", healt)

	routes.SetupRoutes(app)
	app.Listen(":" + s.Port)
}
