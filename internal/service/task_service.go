package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"task_service/internal/model"
	"task_service/internal/repository"
)

type TaskService struct {
	taskRepository repository.TaskRepositoryInterface
	loggerService  LoggerServiceInterface
}

func NewTaskService(taskRepository repository.TaskRepositoryInterface, loggerService LoggerServiceInterface) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
		loggerService:  loggerService,
	}
}

func (t *TaskService) Create(task model.Task) (model.Task, error) {
	if task.Title == "" {
		t.loggerService.AddLog("POST /tasks", false, "Title required!")
		return task, errors.New("title required")
	}
	if task.Status == "" {
		t.loggerService.AddLog("POST /tasks", false, "Status required!")
		return task, errors.New("status required")
	}
	// if !(task.Status != model.StatusPending) && !(task.Status != model.StatusInProcess) && !(task.Status != model.StatusDone) {
	// 	return task, errors.New("status should be one of [pending, inProcess, done]")
	// }

	res, err := t.taskRepository.Create(task)
	if err != nil {
		t.loggerService.AddLog("POST /tasks", false, err.Error())
		return task, err
	}

	t.loggerService.AddLog("POST /tasks", true, fmt.Sprintf("Task with id = %v created successfully", res.ID))

	return res, err
}

func (t *TaskService) GetById(idString string) (model.Task, error) {
	var task model.Task
	id, err := strconv.Atoi(idString)
	if err != nil {
		t.loggerService.AddLog("GET /tasks/"+idString, false, err.Error())
		return task, err
	}
	res, err := t.taskRepository.GetById(id)
	if err != nil {
		t.loggerService.AddLog("GET /tasks/"+idString, false, err.Error())
		return task, err
	}
	t.loggerService.AddLog("GET /tasks"+idString, true, fmt.Sprintf("Task with id = %v received.", idString))
	return res, err
}

func (t *TaskService) GetAll(status string) ([]model.Task, error) {
	res, err := t.taskRepository.GetAll(status)
	if err != nil {
		t.loggerService.AddLog("GET /tasks?status="+status, false, err.Error())
		return res, &model.ServiceError{StatusCode: http.StatusBadRequest, Message: err.Error()}
	}
	var logMessage string
	var logUrl string
	if status == "" {
		logMessage = fmt.Sprintf("Got all tasks with status = %v successfully.", status)
		logUrl = "GET /tasks?status=" + status
	} else {
		logMessage = "Got all tasks successfully."
		logUrl = "GET /tasks"
	}
	t.loggerService.AddLog(logUrl, true, logMessage)
	return res, err
}
