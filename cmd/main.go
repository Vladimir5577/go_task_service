package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task_service/internal/handler"
	"task_service/internal/model"
	"task_service/internal/repository"
	"task_service/internal/service"
	"task_service/internal/storage"
	"time"
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

	serverError := make(chan error, 1)

	go func() {
		log.Println("Server up and running on the port ", servicePort)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server failed:", err)
			serverError <- err
		}
	}()

	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Printf("Server error: %v", err)
	case sig := <-stopChannel:
		log.Printf("Received shutdown signal: %v", sig)
	}

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
		return
	}

	log.Println("Server exited properly")
}
