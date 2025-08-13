package model

type TaskStatus string

const (
	StatusPending   TaskStatus = "Pending"
	StatusInProcess TaskStatus = "InProcess"
	StatusDone      TaskStatus = "Done"
)

type Task struct {
	ID     int        `json:"id"`
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}
