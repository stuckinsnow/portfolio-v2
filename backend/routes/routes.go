package routes

import (
	"portfolio/handlers"
	"portfolio/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	s3 := services.NewS3Service()

	app.Get("/api/portfolio", handlers.PortfolioHandler)
	app.Get("/api/s3/list-objects", s3.HandleListObjects)
	app.Get("/api/s3/object", s3.HandleGetObject)
}
