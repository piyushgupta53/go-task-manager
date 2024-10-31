package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piyushgupta53/task-manager/task"
)

func DeleteTask(c *fiber.Ctx, store *task.TaskStore) error {
	taskID := c.Params("id")

	if !store.Remove(taskID) {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.SendStatus(204)
}
