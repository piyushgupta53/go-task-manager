package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piyushgupta53/task-manager/task"
)

func UpdateTask(c *fiber.Ctx, store *task.TaskStore) error {
	taskID := c.Params("id")
	updatedTask := new(task.Task)

	if err := c.BodyParser(updatedTask); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	updatedTask.ID = taskID
	if !store.Update(updatedTask) {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(updatedTask)
}
