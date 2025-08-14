package service

import (
	"fmt"
	"net/http"
	"strconv"
	"task_service/internal/helper"
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
	err := helper.ValidateRequest(task)
	if err != nil {
		t.loggerService.AddLog("POST /tasks", false, err.Error())
		return task, err
	}

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
	countTasks := len(res)
	var logMessage string
	var logUrl string
	if status != "" {
		logMessage = fmt.Sprintf("Got %v tasks with status = %v successfully.", countTasks, status)
		logUrl = "GET /tasks?status=" + status
	} else {
		logMessage = fmt.Sprintf("Got %v tasks successfully.", countTasks)
		logUrl = "GET /tasks"
	}
	t.loggerService.AddLog(logUrl, true, logMessage)
	return res, err
}
