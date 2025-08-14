package service

import (
	"fmt"
	"task_service/internal/model"
	"time"
)

type LoggerService struct {
	LogChan chan model.Logger
}

func NewLoggerService(logChan chan model.Logger) *LoggerService {
	return &LoggerService{
		LogChan: logChan,
	}
}

func (l *LoggerService) AddLog(action string, success bool, message string) {
	go func() {
		l.LogChan <- model.Logger{
			Timestamp: time.Now(),
			Action:    action,
			Success:   success,
			Message:   message,
		}
	}()
}

func (l *LoggerService) WriteLogging() {
	for log := range l.LogChan {
		fmt.Printf("Time [%s] Action: %s | Success: %v | Message: %s\n", log.Timestamp.Format(time.DateTime), log.Action, log.Success, log.Message)
	}
}
