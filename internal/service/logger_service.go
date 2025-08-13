package service

import (
	"fmt"
	"task_service/internal/model"
	"time"
)

func StartLogging(logChan chan model.Logger) {
	for log := range logChan {
		fmt.Printf("Time [%s] Action: %s, TaskID: %v\n", log.Timestamp.Format(time.RFC3339), log.Action, log.TaskID)
	}
}
