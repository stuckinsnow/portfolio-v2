package routes

import (
	"database/sql"
	"portfolio/handlers"
	"portfolio/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	s3 := services.NewS3Service()

	app.Get("/api/portfolio", handlers.PortfolioHandler)
	app.Get("/api/s3/list-objects", s3.HandleListObjects)
	app.Get("/api/s3/object", s3.HandleGetObject)

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return handlers.GetTodosHandler(c, db)
	})
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		return handlers.PostTodoHandler(c, db)
	})
	app.Put("/api/todos", func(c *fiber.Ctx) error {
		return handlers.PutTodoHandler(c, db)
	})
	app.Delete("/api/todos", func(c *fiber.Ctx) error {
		return handlers.DeleteTodoHandler(c, db)
	})
}
