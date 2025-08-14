package helper

import (
	"errors"
	"fmt"
	"task_service/internal/model"
)

func ValidateRequest(task model.Task) error {
	if task.Title == "" {
		return errors.New("title required")
	}
	if task.Status == "" {
		return errors.New("status required")
	}

	if task.Status != model.StatusPending && task.Status != model.StatusProcess && task.Status != model.StatusDone {
		return fmt.Errorf("status should be one of [%v, %v, %v], got %v", model.StatusPending, model.StatusProcess, model.StatusDone, task.Status)
	}

	return nil
}
