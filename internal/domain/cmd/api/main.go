package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/interfaces/http"
)

func main() {
	// Configurar el modo de Gin según el entorno
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar el router con el middleware de logging
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware())

	// Inicializar el repositorio en memoria con concurrencia
	taskRepo := &repository.TaskRepository{
		Tasks: make(map[string]domain.Task),
		Mutex: sync.Mutex{},
	}

	// Inicializar el servicio de tareas
	taskService := &application.TaskService{
		Repo: taskRepo,
	}

	// Configurar las rutas HTTP
	http.InitializeRoutes(router, taskService)

	// Iniciar el servidor
	serverAddr := ":8080"
	log.Printf("Servidor Gin iniciado en %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}