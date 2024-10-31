package task

import (
	"log"
	"time"
)

var TaskChannel = make(chan *Task, 100)
var resultChannel = make(chan string, 100)

func StartWorkers(workerCount int, store *TaskStore) {
	for i := 0; i < workerCount; i++ {
		go taskWorker(i)
	}

	go monitorResults()
}

func taskWorker(id int) {
	for task := range TaskChannel {
		time.Sleep(2 * time.Second) // Simulate work
		log.Printf("Worker %d processed task: %s\n", id, task.Title)
		task.Status = "completed"
		resultChannel <- task.ID
	}
}

func monitorResults() {
	for taskID := range resultChannel {
		log.Printf("Task completed: %s\n", taskID)
	}
}
