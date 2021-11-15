package routes

import (
	"juanfer2/go-queries/src/modules/queries/controllers"

	"github.com/gofiber/fiber/v2"
)

func skipRoutes(app *fiber.App) {
	app.Get("/queries", controllers.GetQueries)
	app.Post("/queries", controllers.PostCreateQuery)
}

func SetupRoutes(app *fiber.App) {
	skipRoutes(app)
}
