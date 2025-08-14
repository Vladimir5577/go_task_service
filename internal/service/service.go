package service

import "task_service/internal/model"

type TaskServiceInterface interface {
	Create(model.Task) (model.Task, error)
	GetById(string) (model.Task, error)
	GetAll(string) ([]model.Task, error)
}
