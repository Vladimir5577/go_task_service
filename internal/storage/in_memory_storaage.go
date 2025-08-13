package storage

import (
	"sync"
	"task_service/internal/model"
)

type InMemoryStorage struct {
	Tasks  map[int]model.Task
	Mu     sync.RWMutex
	NextId int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Tasks:  make(map[int]model.Task),
		NextId: 1,
	}
}
