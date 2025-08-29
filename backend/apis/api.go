package apis

import (
	"portfolio/services"

	"github.com/gofiber/fiber/v2"
)

type PortfolioData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func SetupRoutes(app *fiber.App) {
	s3 := services.NewS3Service()

	app.Get("/api/portfolio", portfolioHandler)
	app.Get("/api/s3/list-objects", s3.HandleListObjects)
	app.Get("/api/s3/object", s3.HandleGetObject)
}

func portfolioHandler(c *fiber.Ctx) error {
	return c.JSON(PortfolioData{
		Title:       "Full Stack Developer",
		Description: "This is a description",
	})
}
