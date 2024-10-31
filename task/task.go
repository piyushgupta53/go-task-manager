package task

import "time"

type Task struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Priority int       `json:"priority"`
	Status   string    `json:"status"`
	Created  time.Time `json:"created"`
}
