package handlers

import "github.com/gofiber/fiber/v2"

type PortfolioData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func PortfolioHandler(c *fiber.Ctx) error {
	return c.JSON(PortfolioData{
		Title:       "Full Stack Developer",
		Description: "This is a description",
	})
}