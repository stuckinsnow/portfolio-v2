package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetTodosHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT item FROM todos")
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.JSON(fiber.Map{"todos": todos})
}

func PostTodoHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func PutTodoHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func DeleteTodoHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}