package model

import "time"

type Logger struct {
	Timestamp time.Time
	Action    string
	Success   bool
	Message   string
}
