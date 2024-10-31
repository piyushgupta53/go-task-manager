package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piyushgupta53/task-manager/task"
)

func GetTasks(c *fiber.Ctx, store *task.TaskStore) error {
	return c.JSON(store.GetAll())
}
