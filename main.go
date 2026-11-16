package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/app"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/interfaces/http"
)

func main() {
	router := gin.Default()

	// Initialize repository and service
	taskRepo := &repository.TaskRepository{}
	taskService := &application.TaskService{Repo: taskRepo}

	// Initialize routes with the service
	http.InitializeRoutes(router, taskService)

	// Initialize app routes (if any additional routes exist)
	app.InitializeRoutes(router)

	router.Run()
}