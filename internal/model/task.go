package model

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusProcess TaskStatus = "process"
	StatusDone    TaskStatus = "done"
)

type Task struct {
	ID     int        `json:"id"`
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}
