package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	Port string
}

func (server *Server) Init(port string) {
	app := fiber.New()
	app.Use(cors.New())
	server.Port = port

	app.Listen(":" + server.Port)
}
