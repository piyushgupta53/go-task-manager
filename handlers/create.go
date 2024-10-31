package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piyushgupta53/task-manager/task"
)

func CreateTask(c *fiber.Ctx, store *task.TaskStore) error {
	task := new(task.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	store.Add(task)

	// Use the exported TaskChannel

	return c.Status(201).JSON(task)

}
