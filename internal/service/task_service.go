package service

import (
	"errors"
	"strconv"
	"task_service/internal/model"
	"task_service/internal/repository"
	"time"
)

type TaskService struct {
	taskRepository repository.TaskRepositoryInterface
	logChan        chan model.Logger
}

func NewTaskService(taskRepository repository.TaskRepositoryInterface, logChan chan model.Logger) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
		logChan:        logChan,
	}
}

func (t *TaskService) Create(task model.Task) (model.Task, error) {
	if task.Title == "" {
		return task, errors.New("title required")
	}
	if task.Status == "" {
		return task, errors.New("status required")
	}
	// if !(task.Status != model.StatusPending) && !(task.Status != model.StatusInProcess) && !(task.Status != model.StatusDone) {
	// 	return task, errors.New("status should be one of [pending, inProcess, done]")
	// }

	res, err := t.taskRepository.Create(task)

	go func() {
		t.logChan <- model.Logger{
			Timestamp: time.Now(),
			Action:    "POST /tasks",
			TaskID:    res.ID,
		}
	}()

	return res, err
}

func (t *TaskService) GetById(idString string) (model.Task, error) {
	var task model.Task
	id, err := strconv.Atoi(idString)
	if err != nil {
		return task, err
	}
	res, err := t.taskRepository.GetById(id)
	go func() {
		t.logChan <- model.Logger{
			Timestamp: time.Now(),
			Action:    "GET /tasks/" + idString,
			TaskID:    res.ID,
		}
	}()
	return res, err
}

func (t *TaskService) GetAll(status string) ([]model.Task, error) {
	res, err := t.taskRepository.GetAll(status)
	return res, err
}
