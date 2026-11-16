// TODO(reto): archivo stubbeado; implementa los cuerpos.
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/application"
)

type TaskHandler struct {
	taskService *application.TaskService
}

func NewTaskHandler(taskService *application.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
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