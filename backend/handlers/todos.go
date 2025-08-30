package handlers

import (
	"database/sql"
	fmt "fmt"
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

type todo struct {
	Item string
}

func PostTodoHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}

	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT INTO todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Printf("An error occured: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Failed to add todo", "message": err.Error()})
		}
	}
	return c.SendStatus(fiber.StatusOK)
}

func PutTodoHandler(c *fiber.Ctx, db *sql.DB) error {
	oldItem := c.Query("oldItem")
	newItem := c.Query("newItem")
	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newItem, oldItem)
	return c.SendStatus(fiber.StatusOK)
}

func DeleteTodoHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	result, err := db.Exec("DELETE from todos where item=$1", todoToDelete)
	if err != nil {
		log.Printf("An error occured: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete todo", "message": err.Error()})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Todo deleted successfully"})
}
