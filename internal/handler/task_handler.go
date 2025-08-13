package handler

import (
	"net/http"
	"task_service/internal/helper"
	"task_service/internal/model"
	"task_service/internal/service"
)

type TaskHandler struct {
	taskService service.TaskServiceInterface
}

func NewTaskHandler(taskService service.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (t *TaskHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := helper.HandleBody[model.Task](&w, r)
		if err != nil {
			return
		}

		res, err := t.taskService.Create(*task)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}

func (t *TaskHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		res, err := t.taskService.GetById(idString)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}

func (t *TaskHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		res, err := t.taskService.GetAll(status)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}
