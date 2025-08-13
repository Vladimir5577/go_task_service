package repository

import (
	"fmt"
	"task_service/internal/model"
	"task_service/internal/storage"
)

type TaskRepositoryInterface interface {
	Create(model.Task) (model.Task, error)
	GetById(int) (model.Task, error)
	GetAll(string) ([]model.Task, error)
}

type TaskRepository struct {
	storage *storage.InMemoryStorage
}

func NewTaskRepository(storage *storage.InMemoryStorage) *TaskRepository {
	return &TaskRepository{
		storage: storage,
	}
}

func (t *TaskRepository) Create(task model.Task) (model.Task, error) {
	t.storage.Mu.Lock()
	defer t.storage.Mu.Unlock()

	task.ID = t.storage.NextId
	t.storage.Tasks[t.storage.NextId] = task
	t.storage.NextId++

	return task, nil
}

func (t *TaskRepository) GetById(id int) (model.Task, error) {
	t.storage.Mu.RLock()
	defer t.storage.Mu.RUnlock()

	task, found := t.storage.Tasks[id]
	if !found {
		return task, fmt.Errorf("task with id = %v not found", id)
	}
	return task, nil
}

func (t *TaskRepository) GetAll(status string) ([]model.Task, error) {
	t.storage.Mu.RLock()
	defer t.storage.Mu.RUnlock()

	var tasks []model.Task
	for _, task := range t.storage.Tasks {
		if status == "" || task.Status == model.TaskStatus(status) {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}
