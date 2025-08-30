package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"portfolio/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.SetupRoutes(app, db)

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
