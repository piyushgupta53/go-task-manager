package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/piyushgupta53/task-manager/handlers"
	"github.com/piyushgupta53/task-manager/task"
)

func main() {
	app := fiber.New()

	// Initialize TaskStore and start workers
	taskStore := task.NewTaskStore()
	task.StartWorkers(3, taskStore) // Start 3 workers to process tasks

	// Routes
	app.Post("/tasks", func(c *fiber.Ctx) error { return handlers.CreateTask(c, taskStore) })
	app.Get("/tasks", func(c *fiber.Ctx) error { return handlers.GetTasks(c, taskStore) })
	app.Put("/tasks/:id", func(c *fiber.Ctx) error { return handlers.UpdateTask(c, taskStore) })
	app.Delete("/tasks/:id", func(c *fiber.Ctx) error { return handlers.DeleteTask(c, taskStore) })

	log.Fatal(app.Listen(":3000"))
}
