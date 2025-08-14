package repository

import "task_service/internal/model"

type TaskRepositoryInterface interface {
	Create(model.Task) (model.Task, error)
	GetById(int) (model.Task, error)
	GetAll(string) ([]model.Task, error)
}
