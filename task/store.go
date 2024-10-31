package task

import (
	"time"

	"github.com/google/uuid"
)

type TaskStore struct {
	tasks        []*Task
	addChan      chan *Task
	removeChan   chan string
	updateChan   chan *Task
	getAllChan   chan bool
	responseChan chan interface{}
}

func NewTaskStore() *TaskStore {
	store := &TaskStore{
		tasks:        make([]*Task, 0),
		addChan:      make(chan *Task),
		removeChan:   make(chan string),
		updateChan:   make(chan *Task),
		getAllChan:   make(chan bool),
		responseChan: make(chan interface{}),
	}
	go store.run()
	return store
}

func (ts *TaskStore) run() {
	for {
		select {
		case task := <-ts.addChan:
			ts.addTask(task)
			ts.responseChan <- true
		case id := <-ts.removeChan:
			success := ts.removeTask(id)
			ts.responseChan <- success
		case task := <-ts.updateChan:
			success := ts.updateTask(task)
			ts.responseChan <- success
		case <-ts.getAllChan:
			ts.responseChan <- ts.tasks
		}
	}
}

func (ts *TaskStore) addTask(task *Task) {
	ts.tasks = append(ts.tasks, task)

	for i := len(ts.tasks) - 1; i > 0; i-- {
		if ts.tasks[i].Priority > ts.tasks[i-1].Priority {
			ts.tasks[i], ts.tasks[i-1] = ts.tasks[i-1], ts.tasks[i]
		}
	}
}

func (ts *TaskStore) removeTask(id string) bool {
	for i, task := range ts.tasks {
		if task.ID == id {
			ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (ts *TaskStore) updateTask(updatedTask *Task) bool {
	for i, task := range ts.tasks {
		if task.ID == updatedTask.ID {
			ts.tasks[i] = updatedTask
			return true
		}
	}
	return false
}

func (ts *TaskStore) Add(task *Task) {
	task.ID = uuid.New().String()
	task.Created = time.Now()
	task.Status = "Pending"

	ts.addChan <- task
	<-ts.responseChan
}

func (ts *TaskStore) Remove(id string) bool {
	ts.removeChan <- id
	response := <-ts.responseChan

	return response.(bool)
}

func (ts *TaskStore) Update(task *Task) bool {
	ts.updateChan <- task
	response := <-ts.responseChan

	return response.(bool)
}

func (ts *TaskStore) GetAll() []*Task {
	ts.getAllChan <- true
	response := <-ts.responseChan

	return response.([]*Task)
}
