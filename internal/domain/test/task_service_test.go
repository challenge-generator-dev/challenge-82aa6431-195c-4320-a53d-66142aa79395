// TODO(reto): archivo stubbeado; implementa las pruebas unitarias.
package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/repository"
)

func TestCreateTask_Success(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para creación exitosa de tarea")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, err := service.CreateTask("Nueva tarea de prueba")

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "Nueva tarea de prueba", task.Title)
	assert.Equal(t, domain.TaskStatusPending, task.Status)
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para creación con título vacío")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, err := service.CreateTask("")

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.Equal(t, domain.ErrTitleEmpty, err)
}

func TestCreateTask_TitleTooLong(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para título que excede 100 caracteres")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	longTitle := string(make([]byte, 101))
	for i := range longTitle {
		longTitle = string(rune('a' + i%26))
	}

	task, err := service.CreateTask(longTitle)

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.Equal(t, domain.ErrTitleTooLong, err)
}

func TestGetAllTasks_Empty(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para listar tareas cuando no hay tareas")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	tasks, err := service.GetAllTasks()

	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.Empty(t, tasks)
}

func TestGetAllTasks_WithTasks(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para listar tareas existentes")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	_, _ = service.CreateTask("Tarea 1")
	_, _ = service.CreateTask("Tarea 2")

	tasks, err := service.GetAllTasks()

	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.Len(t, tasks, 2)
}

func TestUpdateTaskStatus_Success(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para actualización exitosa de estado")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, _ := service.CreateTask("Tarea para actualizar")
	updatedTask, err := service.UpdateTaskStatus(task.ID, domain.TaskStatusCompleted)

	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, domain.TaskStatusCompleted, updatedTask.Status)
}

func TestUpdateTaskStatus_InvalidID(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para actualización con ID inválido")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	updatedTask, err := service.UpdateTaskStatus("id-invalido", domain.TaskStatusCompleted)

	assert.Error(t, err)
	assert.Nil(t, updatedTask)
	assert.Equal(t, domain.ErrTaskNotFound, err)
}

func TestUpdateTaskStatus_InvalidStatus(t *testing.T) {
	t.Skip("TODO(reto): implementar prueba unitaria para actualización con estado inválido")

	repo := &repository.TaskRepository{}
	service := &application.TaskService{Repo: repo}

	task, _ := service.CreateTask("Tarea con estado inválido")
	invalidStatus := domain.TaskStatus("invalid_status")

	updatedTask, err := service.UpdateTaskStatus(task.ID, invalidStatus)

	assert.Error(t, err)
	assert.Nil(t, updatedTask)
	assert.Equal(t, domain.ErrInvalidTaskStatus, err)
}