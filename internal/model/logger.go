package model

import "time"

type Logger struct {
	Timestamp time.Time
	Action    string
	TaskID    int
}
