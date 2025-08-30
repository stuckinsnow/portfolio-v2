package main

import (
	"fmt"
	"log"
	"os"
	"portfolio/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Static("/", "../dist")

	app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("../dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
