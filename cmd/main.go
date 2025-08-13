package main

import (
	"log"
	"net/http"
	"task_service/internal/handler"
	"task_service/internal/model"
	"task_service/internal/repository"
	"task_service/internal/service"
	"task_service/internal/storage"
)

const (
	servicePort = ":8080"
)

func main() {
	logChan := make(chan model.Logger)
	go service.StartLogging(logChan)

	mux := http.NewServeMux()

	storage := storage.NewInMemoryStorage()
	taskRepository := repository.NewTaskRepository(storage)
	taskService := service.NewTaskService(taskRepository, logChan)
	taskHandler := handler.NewTaskHandler(taskService)

	mux.HandleFunc("POST /tasks", taskHandler.Create())
	mux.HandleFunc("GET /tasks", taskHandler.GetAll())
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetById())

	server := &http.Server{
		Addr:    servicePort,
		Handler: mux,
	}

	log.Println("Server up and running on the port ", servicePort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed:", err)
	}
}
