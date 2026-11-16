// TODO(reto): archivo stubbeado; implementa los cuerpos.
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
)

func InitializeRoutes(router *gin.Engine, taskService *application.TaskService) {
	taskHandler := NewTaskHandler(taskService)

	// Mount routes according to the contract
	router.POST("/tasks", taskHandler.HandleCreateTask)
	router.GET("/tasks", taskHandler.HandleGetAllTasks)
	router.PUT("/tasks/:id", taskHandler.HandleUpdateTaskStatus)
}

// TaskHandler defines the HTTP handlers for task operations.
type TaskHandler struct {
	Service *application.TaskService
}

func NewTaskHandler(taskService *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: taskService,
	}
}

func (h *TaskHandler) HandleCreateTask(c *gin.Context) {
	panic("TODO(reto): HandleCreateTask")
}

func (h *TaskHandler) HandleGetAllTasks(c *gin.Context) {
	panic("TODO(reto): HandleGetAllTasks")
}

func (h *TaskHandler) HandleUpdateTaskStatus(c *gin.Context) {
	panic("TODO(reto): HandleUpdateTaskStatus")
}